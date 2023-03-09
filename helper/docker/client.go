package docker

import (
	"github.com/docker/docker/client"
)

type DockerClient struct {
	*client.Client
}

func New(ops ...client.Opt) (*DockerClient, error) {

	ops = append(ops, client.WithAPIVersionNegotiation())
	client, err := client.NewClientWithOpts(ops...)
	return &DockerClient{client}, err

}
