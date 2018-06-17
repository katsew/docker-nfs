package volume

import (
	"fmt"
	"os/exec"
)

func NewDockerVolumeRmCommand(name string) func() error {
	return func() error {
		rmCommand := fmt.Sprintf(`docker volume rm %s`, name)
		cmd := exec.Command("/bin/sh", "-c", rmCommand)
		err := cmd.Run()
		if err != nil {
			return err
		}
		return nil
	}
}
