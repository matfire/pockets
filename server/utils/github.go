package utils

import (
	"encoding/json"
	"io"
	"net/http"

	sharedv1 "github.com/matfire/pockets/shared/v1"
)

type PBRelease struct {
	Name string `json:"name"`
	Tag  string `json:"tag_name"`
}

func GetVersions() (*sharedv1.GetVersionsResponse, error) {
	var data []PBRelease
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/repos/pocketbase/pocketbase/releases", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Add("Accept", "application/vnd.github+json")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	json.Unmarshal(b, &data)
	var data_connect []*sharedv1.Version
	for _, version := range data {
		data_connect = append(
			data_connect,
			&sharedv1.Version{
				Name: version.Name,
				Tag:  version.Tag,
			},
		)
	}
	res_connect := &sharedv1.GetVersionsResponse{
		Versions: data_connect,
	}
	return res_connect, nil
}
