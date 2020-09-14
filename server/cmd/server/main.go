package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/container-tools/boxit/api"
	"github.com/container-tools/boxit/server/pkg/builder"
	builderapi "github.com/container-tools/boxit/server/pkg/builder/api"
)

func main() {
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

	response, err := builder.Default(img)
	if err != nil {
		if err.Error() == builderapi.BuildFailure {
			log.Printf("Error while processing request: %v", err)
			res.WriteHeader(400)
		} else {
			log.Printf("Error while creating response: %v", err)
			res.WriteHeader(500)
		}

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
