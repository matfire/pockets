package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/matfire/pockets/server/embeds"
	sharedv1 "github.com/matfire/pockets/shared/v1"
)

func CheckImage(data *sharedv1.CheckImageRequest) (*sharedv1.CheckImageResponse, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	images, err := cli.ImageList(context.Background(), image.ListOptions{
		All: true,
	})
	if err != nil {
		return nil, err
	}
	imgIdx := slices.IndexFunc(images, func(el image.Summary) bool {
		repoIdx := slices.IndexFunc(el.RepoTags, func(t string) bool {
			return t == fmt.Sprintf("pockets:%s", data.Version)
		})
		return repoIdx != -1
	})

	response := sharedv1.CheckImageResponse{}

	if imgIdx == -1 {
		response.Exists = false
	} else {
		response.Exists = true
	}
	return &response, nil
}

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
