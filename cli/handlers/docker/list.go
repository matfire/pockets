package docker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/matfire/pockets/cli/config"
)

type ListResponse struct {
	Message string `json:"Message"`
}

func List(config *config.App) {
	fmt.Print("listing containers")
	res, err := http.Get(fmt.Sprintf("%s/v1/status", config.Endpoint))
	if err != nil {
		panic("could not get data")
	}
	b, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	var data ListResponse
	json.Unmarshal(b, &data)
	fmt.Printf("message was %s", data.Message)

}
