package docker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/matfire/pockets/cli/config"
)

type CreateRequestBody struct {
	Name string
}

func Create(config *config.App, name string) {
	body := CreateRequestBody{Name: name}
	b, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("could not marshal body in create request \n")
	}
	res, err := http.Post(fmt.Sprintf("%s/v1/create", config.Endpoint), "application/json", bytes.NewBuffer(b))
	if err != nil {
		fmt.Printf("create request failed with error %v", err)
		return
	}
	fmt.Printf("create response: %v", res)
}
