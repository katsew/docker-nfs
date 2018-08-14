package volume

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/katsew/docker-nfs/pkg/common"
)

type volumeOptions struct {
	Device string
	O      string
	Type   string
}

type inspectResult struct {
	CreatedAt  string
	Driver     string
	Labels     interface{}
	Mountpoint string
	Name       string
	Options    volumeOptions
	Scope      string
}

var inspectResults []inspectResult

func NewDockerVolumeInspectCommand(name string) func() error {
	return func() error {
		inspectCmd := fmt.Sprintf(`docker volume inspect %s`, name)
		cmd := exec.Command("/bin/sh", "-c", inspectCmd)
		b, err := cmd.Output()
		err = json.Unmarshal(b, &inspectResults)
		if err != nil {
			return err
		}
		if len(inspectResults) > 0 {
			return common.DockerVolumeAlreadyExists{
				Name:   inspectResults[0].Name,
				Device: inspectResults[0].Options.Device,
				Option: inspectResults[0].Options.O,
				Err:    common.ErrVolumeAlreadyExist,
			}
		}
		return nil
	}
}
