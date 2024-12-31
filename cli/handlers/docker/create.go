package docker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/charmbracelet/huh/spinner"
	"github.com/matfire/pockets/cli/config"
)

type CreateRequestBody struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type ImageCreateRequestBody struct {
	Version string `json:"version"`
}

type ImageCheckResponse struct {
	Exists bool `json:"exists"`
}

func Create(config *config.App, name string, version string) {
	imageExists := false
	checkImage := func() {
		res, err := http.Get(fmt.Sprintf("%s/v1/image/check/%s", config.Endpoint, version))
		var data ImageCheckResponse
		if err != nil {
			panic(err)
		}
		b, err := io.ReadAll(res.Body)
		json.Unmarshal(b, &data)
		defer res.Body.Close()
		imageExists = data.Exists
	}

	createImage := func() {
		body := ImageCreateRequestBody{Version: version}
		b, err := json.Marshal(body)
		if err != nil {
			panic("could not marshal image create request body")
		}
		_, err = http.Post(fmt.Sprintf("%s/v1/image/new", config.Endpoint), "application/json", bytes.NewBuffer(b))
	}

	//TODO customize retries & interval
	waitForImage := func() {
		iterations := 1
		for {
			if iterations > 5 {
				panic("image is taking too long")
			}
			res, err := http.Get(fmt.Sprintf("%s/v1/image/check/%s", config.Endpoint, version))
			if err != nil {
				break
			}
			var data ImageCheckResponse
			b, err := io.ReadAll(res.Body)
			json.Unmarshal(b, &data)
			defer res.Body.Close()
			if data.Exists {
				break
			}
			iterations++
			time.Sleep(5 * time.Second)
		}
	}

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
	if err := spinner.New().Title("Checking for image existence...").Action(checkImage).Run(); err != nil {
		fmt.Println(err)
	}
	if imageExists {
		if err := spinner.New().Title("Creating Container...").Action(createContainer).Run(); err != nil {
			fmt.Println(err)
		}
	} else {
		if err := spinner.New().Title("Creating Image...").Action(createImage).Run(); err != nil {
			fmt.Println(err)
		}
		if err := spinner.New().Title("Waiting For Image...").Action(waitForImage).Run(); err != nil {
			fmt.Println(err)
		}
		if err := spinner.New().Title("Creating Container...").Action(createContainer).Run(); err != nil {
			fmt.Println(err)
		}
	}

}
