package kubernetes

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	v1 "github.com/apache/camel-k/pkg/apis/camel/v1"
	camelclient "github.com/apache/camel-k/pkg/client/camel/clientset/versioned"
	"github.com/container-tools/boxit/api"
	builderapi "github.com/container-tools/boxit/server/pkg/builder/api"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var namespace string
var client *camelclient.Clientset

func init() {
	namespace = os.Getenv("NAMESPACE")
	if namespace == "" {
		fmt.Println(`No "NAMESPACE" environment variable present`)
		os.Exit(1)
	}

	client = camelclient.NewForConfigOrDie(GetKubernetesConfigOrDie())
}

func Builder(img api.ImageRequest) (api.ImageResult, error) {
	_, err := findOrCreateKit(img)
	if err != nil {
		log.Printf("Error while creating resources: %v", err)
		return api.ImageResult{}, err
	}

	err = waitForCompletion(img)
	if err != nil {
		log.Printf("Error during build creation: %v", err)
		return api.ImageResult{}, err
	}

	return buildResponse(img)
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
			return errors.New(builderapi.BuildFailure)
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
