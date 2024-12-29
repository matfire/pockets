package docker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/matfire/pockets/cli/config"
)

type Container struct {
	Name string `json:"name"`
}

type ListResponse struct {
	Containers []Container `json:"containers"`
}

func List(config *config.App) {
	fmt.Print("listing containers")

	var data ListResponse
	getContainers := func() {
		res, err := http.Get(fmt.Sprintf("%s/v1/status", config.Endpoint))
		if err != nil {
			panic("could not get data")
		}
		b, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		json.Unmarshal(b, &data)

	}
	if err := spinner.New().Title("Fetching containers...").Action(getContainers).Run(); err != nil {
		fmt.Println(err)
	}
	t := table.New().Headers("Name")

	for _, container := range data.Containers {
		t.Row(container.Name)
	}
	fmt.Println(t)
}
