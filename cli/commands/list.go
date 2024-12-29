package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

type ListResponse struct {
	Message string `json:"Message"`
}

func CreateListCommand() *cobra.Command {
	var listCommand = &cobra.Command{
		Use:     "list",
		Example: "pocketsctl list",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("listing containers")
			res, err := http.Get("http://127.0.0.1:3000/v1/status")
			if err != nil {
				panic("could not get data")
			}
			b, err := io.ReadAll(res.Body)
			defer res.Body.Close()
			var data ListResponse
			json.Unmarshal(b, &data)
			fmt.Printf("message was %s", data.Message)

		},
	}
	return listCommand
}
