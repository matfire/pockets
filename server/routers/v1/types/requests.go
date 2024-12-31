package types

type ContainerCreateBody struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type ImageCreateBody struct {
	Version string `json:"version"`
}
