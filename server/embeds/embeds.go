package embeds

import _ "embed"

//go:embed Dockerfile
var Dockerfile string

//go:embed Dockerfile
var DockerfileBytes []byte
