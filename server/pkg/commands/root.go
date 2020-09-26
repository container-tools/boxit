package commands

import (
	"encoding/json"
	"fmt"
	"github.com/container-tools/boxit/api"
	"github.com/container-tools/boxit/server/pkg/builder"
	builderapi "github.com/container-tools/boxit/server/pkg/builder/api"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
)

type rootOptions struct {
	builderapi.BuilderOptions
	host string
	port int
}

func NewCmdRoot() *cobra.Command {
	options := rootOptions{}

	cmd := cobra.Command{
		Use:   "boxit-server",
		Short: "Runs the boxit server",
		Run: func(cmd *cobra.Command, args []string) {
			start(options)
		},
	}

	cmd.Flags().StringVar(&options.host, "host", "0.0.0.0", `The server host`)
	cmd.Flags().IntVarP(&options.port, "port", "p", 8080, `The server port`)

	cmd.Flags().StringVar(&options.Registry, "registry", "localhost:5000", `The container image registry used to host images (e.g. "docker.io/yourid")`)
	cmd.Flags().BoolVar(&options.Insecure, "insecure", true, `If the container image registry uses an insecure protocol`)

	return &cmd
}

func start(options rootOptions) {
	http.HandleFunc("/images", func(res http.ResponseWriter, req *http.Request) {
		require(options, res, req)
	})

	log.Printf("Using target registry %s (insecure=%v)\n", options.Registry, options.Insecure)
	log.Printf("Server listening on %s:%d...\n", options.host, options.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", options.host, options.port), nil))
}

func require(options rootOptions, res http.ResponseWriter, req *http.Request) {
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

	response, err := builder.Default(options.BuilderOptions, img)
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
