package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/matfire/pockets/cli/config"
	"github.com/matfire/pockets/cli/handlers/docker"
	"github.com/spf13/cobra"
)

type PBRelease struct {
	Name    string `json:"name"`
	TagName string `json:"tag_name"`
}

func CreateCreateCommand(config *config.App) *cobra.Command {
	var createCommand = &cobra.Command{
		Use:     "create",
		Example: "pocketsctl create test-1",
		Args:    cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var data []PBRelease

			var selectedVersion string

			fetchVersions := func() {
				req, err := http.NewRequest(http.MethodGet, "https://api.github.com/repos/pocketbase/pocketbase/releases", nil)
				if err != nil {
					panic("could not create http get request")
				}
				req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
				req.Header.Add("Accept", "application/vnd.github+json")
				client := http.Client{}
				res, err := client.Do(req)
				if err != nil {
					panic(err)
				}
				b, err := io.ReadAll(res.Body)
				defer res.Body.Close()
				json.Unmarshal(b, &data)
			}

			var name string
			var version string
			if len(args) > 0 {
				name = args[0]
			} else {
				huh.NewInput().Title("what's the name of the container?").Value(&name).Run()
			}
			err := spinner.New().
				Title("Fetching available versions...").
				Action(fetchVersions).
				Run()
			fetchVersions()
			var options []huh.Option[string]
			for _, v := range data {
				options = append(options, huh.NewOption(v.Name, v.TagName))
			}
			version, err = cmd.Flags().GetString("version")
			if err != nil {
				panic(err)
			}
			if version == "" {
				huh.NewSelect[string]().Title("Which version would you like to deploy?").Options(options...).Value(&version).Run()
				fmt.Printf("selected version is %s\n", selectedVersion)
			} else {
				idx := slices.IndexFunc(data, func(el PBRelease) bool {
					return el.TagName == version
				})
				if idx == -1 {
					fmt.Printf("Specified version %s is not valid, aborting...", version)
					return
				}
			}
			docker.Create(config, name, version)
		},
	}
	createCommand.Flags().StringP("version", "v", "", "specify a version to deploy")
	return createCommand
}
