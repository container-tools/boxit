package commands

import (
	"fmt"
	"github.com/container-tools/boxit/api"
	"github.com/container-tools/boxit/client"
	"github.com/spf13/cobra"
)

type createOptions struct {
	Platform     string
	Dependencies []string
}

func NewCmdCreate() *cobra.Command {
	options := createOptions{}

	cmd := cobra.Command{
		Use:   "create",
		Short: "Creates or returns the address of a boxit image",
		RunE: func(cmd *cobra.Command, args []string) error {
			c := client.New()
			img := api.Image{
				Platform: api.Platform(options.Platform),
			}
			for _, d := range options.Dependencies {
				img.Dependencies = append(img.Dependencies, api.Dependency(d))
			}
			image, err := c.Create(img)
			if err != nil {
				return err
			}
			fmt.Println(image)
			return nil
		},
	}

	cmd.Flags().StringArrayVarP(&options.Dependencies, "dependency", "d", nil, `A list of dependencies to add in full format (e.g. mvn:org.apache.camel:camel-core:3.3.0)`)
	cmd.Flags().StringVarP(&options.Platform, "platform", "p", string(api.PlatformJVM), `The target platform for the image`)

	return &cmd
}
