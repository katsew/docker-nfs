package volume

import (
	"fmt"
	"os/exec"
)

func NewDockerVolumeCreateCommand(cwd string, name string) func() error {
	return func() error {
		createCommand := fmt.Sprintf(`docker volume create \
		--driver local \
		--opt type=nfs \
		--opt o=addr=host.docker.internal,rw,vers=3,tcp,fsc,actimeo=2 \
		--opt device=:%s \
		%s
`, cwd, name)
		cmd := exec.Command("/bin/sh", "-c", createCommand)
		err := cmd.Run()
		if err != nil {
			return err
		}
		return nil

	}
}
