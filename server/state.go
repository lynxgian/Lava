package server

import (
	"context"
	"errors"
	lxdUtil "github.com/Hye-Ararat/Lava/lxd"
	"github.com/docker/docker/client"
	lxd "github.com/lxc/lxd/client"
)

func GetState(server string) (string, error) {
	serverType, err := GetType(server)
	if err != nil {
		return "", errors.New("could not identify server type")
	}
	switch serverType {
	case "kvm", "n-vps":
		c, err := lxd.ConnectLXDUnix("/var/snap/lxd/common/lxd/unix.socket", nil)
		if err != nil {
			return "", err
		}
		state, _, err := c.GetInstanceState(lxdUtil.ConvertId(server))
		if err != nil {
			return "", err
		}
		return state.Status, nil
	case "docker":
		cli, err := client.NewClientWithOpts()
		if err != nil {
			return "", err
		}
		ctx := context.Background()
		container, err := cli.ContainerInspect(ctx, server)
		if err != nil {
			return "", err
		}
		return container.State.Status, nil
	default:
		return "", nil
	}
}
