package main

import (
	"encoding/json"
	"errors"
	"fmt"
	v1 "github.com/apache/camel-k/pkg/apis/camel/v1"
	"github.com/container-tools/boxit/api"
	"github.com/container-tools/boxit/server/kubernetes"
	"io/ioutil"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"net/http"
	"os"
	"time"

	camelclient "github.com/apache/camel-k/pkg/client/camel/clientset/versioned"
)

var namespace string
var client *camelclient.Clientset

const buildFailure = "image build failed"

func main() {
	namespace = os.Getenv("NAMESPACE")
	if namespace == "" {
		fmt.Println(`No "NAMESPACE" environment variable present`)
		os.Exit(1)
	}

	client = camelclient.NewForConfigOrDie(kubernetes.GetKubernetesConfigOrDie())

	http.HandleFunc("/images", require)

	port := 8080
	log.Printf("Server listening on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func require(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		res.WriteHeader(405)
		return
	}

	content, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("An error occurred: %v", err)
		res.WriteHeader(400)
		return
	}
	var img api.ImageRequest
	if err := json.Unmarshal(content, &img); err != nil {
		log.Printf("An error occurred during unmarshal: %v", err)
		res.WriteHeader(400)
		return
	}

	if img.Platform != api.PlatformJVM {
		log.Printf("Unsupported platform: %q", img.Platform)
		res.WriteHeader(400)
		return
	}

	_, err = findOrCreateKit(img)
	if err != nil {
		log.Printf("Error while creating resources: %v", err)
		res.WriteHeader(500)
		return
	}

	err = waitForCompletion(img)
	if err != nil {
		log.Printf("Error during build creation: %v", err)
		if err.Error() == buildFailure {
			res.WriteHeader(400)
		} else {
			res.WriteHeader(500)
		}
		return
	}

	response, err := buildResponse(img)
	if err != nil {
		log.Printf("Error while creating response: %v", err)
		res.WriteHeader(500)
		return
	}
	resData, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error while marshalling response: %v", err)
		res.WriteHeader(500)
		return
	}
	res.Header().Add("Content-Type", "application/json")
	res.Write(resData)
}

func findOrCreateKit(img api.ImageRequest) (*v1.IntegrationKit, error) {
	kit, err := findKit(img)
	if err != nil && k8serrors.IsNotFound(err) {
		return createKit(img)
	}
	return kit, err
}

func findKit(img api.ImageRequest) (*v1.IntegrationKit, error) {
	name := img.Hash()
	return client.CamelV1().IntegrationKits(namespace).Get(name, metav1.GetOptions{})
}

func createKit(img api.ImageRequest) (*v1.IntegrationKit, error) {
	name := img.Hash()
	kit := v1.IntegrationKit{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
	}

	for _, d := range img.Dependencies {
		kit.Spec.Dependencies = append(kit.Spec.Dependencies, string(d))
	}

	return client.CamelV1().IntegrationKits(namespace).Create(&kit)
}

func waitForCompletion(img api.ImageRequest) error {
	// Waiting for the final image is only necessary because the tag is non-deterministic in the Camel K builder
	// Technically we can return the final image before building it
	for i := 0; i < 600; i++ {
		kit, err := findKit(img)
		if err != nil {
			return err
		}
		if kit.Status.Phase == v1.IntegrationKitPhaseError {
			return errors.New(buildFailure)
		}
		if kit.Status.Image == "" {
			time.Sleep(1 * time.Second)
			continue
		}
		return nil
	}
	return errors.New("image build is not finished yet")
}

func buildResponse(img api.ImageRequest) (api.ImageResult, error) {
	kit, err := findKit(img)
	if err != nil {
		return api.ImageResult{}, err
	}
	res := api.ImageResult{
		ID: kit.Status.Image,
	}
	for _, a := range kit.Status.Artifacts {
		resA := api.Artifact{
			ID:       a.ID,
			Checksum: a.Checksum,
			Target:   a.Target,
			Location: a.Location,
		}
		res.Artifacts = append(res.Artifacts, resA)
	}
	return res, nil
}
