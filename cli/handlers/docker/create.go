package docker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charmbracelet/huh/spinner"
	"github.com/matfire/pockets/cli/config"
)

type CreateRequestBody struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func Create(config *config.App, name string, version string) {

	createContainer := func() {
		body := CreateRequestBody{Name: name, Version: version}
		b, err := json.Marshal(body)
		if err != nil {
			fmt.Printf("could not marshal body in create request \n")
		}
		_, err = http.Post(fmt.Sprintf("%s/v1/create", config.Endpoint), "application/json", bytes.NewBuffer(b))
		if err != nil {
			fmt.Printf("create request failed with error %v", err)
			return
		}
	}
	if err := spinner.New().Title("Creating Container...").Action(createContainer).Run(); err != nil {
		fmt.Println(err)
	}

}
