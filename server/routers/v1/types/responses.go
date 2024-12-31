package types

type StatusResponse struct {
	Message string
}

type Container struct {
	Name string `json:"name"`
}

type ListResponse struct {
	Containers []Container `json:"containers"`
}

type ImageCheckResponse struct {
	Exists bool `json:"exists"`
}
