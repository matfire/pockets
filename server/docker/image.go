package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/charmbracelet/log"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
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
		All:     true,
		Filters: filters.NewArgs(filters.Arg("reference", fmt.Sprintf("pockets:%s", data.Version))),
	})
	if err != nil {
		return nil, err
	}
	response := sharedv1.CheckImageResponse{}

	if len(images) == 0 {
		response.Exists = false
	} else {
		response.Exists = true
	}
	return &response, nil
}

func CreateImage(data *sharedv1.CreateImageRequest) (*sharedv1.CreateImageResponse, error) {
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
	if imgIdx == -1 {
		err = CreatePBImage(data.Version)
		if err != nil {
			return nil, err
		}
	}
	return &sharedv1.CreateImageResponse{}, nil
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
	tries := 0
	for {
		if tries > 5 {
			log.Error("Got to end of waiting loop; this should not happen")
			break
		}
		//TODO this should be configurable
		time.Sleep(2 * time.Second)
		info, _ := CheckImage(&sharedv1.CheckImageRequest{Version: tag})
		if info.Exists {
			break
		} else {
			tries++
		}
	}
	return err
}
