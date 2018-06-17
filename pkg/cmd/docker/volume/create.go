package volume

import (
	"fmt"
	"os/exec"
)

const (
	DefaultNfsOptions = "addr=host.docker.internal,rw"
)

func NewDockerVolumeCreateCommand(cwd string, name string, o string) func() error {

	if o == "" {
		o = DefaultNfsOptions
	}
	return func() error {
		createCommand := fmt.Sprintf(`docker volume create \
		--driver local \
		--opt type=nfs \
		--opt o=%s \
		--opt device=:%s \
		%s
`, o, cwd, name)
		cmd := exec.Command("/bin/sh", "-c", createCommand)
		err := cmd.Run()
		if err != nil {
			return err
		}
		return nil

	}
}
