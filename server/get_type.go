package server

import (
	"context"
	"errors"
	"github.com/docker/docker/client"
	lxd "github.com/lxc/lxd/client"
)
import lxdUtil "github.com/Hye-Ararat/Lava/lxd"

func GetType(server string) (string, error) {
	c, err := lxd.ConnectLXDUnix("/var/snap/lxd/common/lxd/unix.socket", nil)
	if err != nil {
		return "An Error Occurred", err
	}
	instance, _, lxdError := c.GetInstance(lxdUtil.ConvertId(server))
	if lxdError == nil {
		if instance.Type == "container" {
			return "n-vps", nil
		}
		if instance.Type == "virtual-machine" {
			return "kvm", nil
		}
	}
	cli, err := client.NewClientWithOpts()
	ctx := context.Background()
	_, errDocker := cli.ContainerInspect(ctx, server)
	if errDocker == nil {
		return "docker", nil
	}
	return "", errors.New("could not identify server type")

}
