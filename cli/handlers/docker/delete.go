package docker

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/huh/spinner"
	"github.com/matfire/pockets/cli/config"
)

func Delete(config *config.App, name string) {
	var err error
	stopContainer := func() {
		request, requestError := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/v1/delete/%s", config.Endpoint, name), nil)
		if requestError != nil {
			fmt.Printf("create request failed with error %v", requestError)
			return
		}
		client := &http.Client{}
		_, requestError = client.Do(request)
		if requestError != nil {
			fmt.Printf("create request failed with error %v", requestError)
			return
		}
		err = requestError
	}
	if spinnerErr := spinner.New().Title("Deleting Container...").Action(stopContainer).Run(); spinnerErr != nil {
		fmt.Println(spinnerErr)
	}
	if err != nil {
		fmt.Printf("create request failed with error %v", err)
		return
	}

}
