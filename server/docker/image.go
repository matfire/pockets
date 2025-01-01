package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/matfire/pockets/server/embeds"
)

func CreatePBImage(tag string) error {
	var err error

	cli, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()
	dockerFile := "Dockerfile"
	readDockerFile := embeds.DockerfileBytes
	tarHeader := &tar.Header{
		Name: dockerFile,
		Size: int64(len(readDockerFile)),
	}
	err = tw.WriteHeader(tarHeader)
	if err != nil {
		log.Fatal(err, " :unable to write tar header")
		return err
	}
	_, err = tw.Write(readDockerFile)
	if err != nil {
		log.Fatal(err, " :unable to write tar body")
		return err
	}
	dockerFileTarReader := bytes.NewReader(buf.Bytes())

	tags := []string{fmt.Sprintf("pockets:%s", tag)}
	versionArg := strings.Split(tag, "v")[1]
	_, err = cli.ImageBuild(context.Background(), dockerFileTarReader, types.ImageBuildOptions{
		Context:    dockerFileTarReader,
		Dockerfile: dockerFile,
		Tags:       tags,
		BuildArgs:  map[string]*string{"PB_VERSION": &versionArg},
		Remove:     true,
	})
	return err
}
