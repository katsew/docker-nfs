package cmd

import (
	"log"
	"os"

	"strings"

	volumeCmd "github.com/katsew/docker-nfs/pkg/cmd/docker/volume"
	"github.com/katsew/docker-nfs/pkg/common"
	"github.com/spf13/cobra"
)

var (
	forceRecreate bool
	nfsOptions    string
	name          string
)

func init() {
	CreateVolumeCommand.Flags().BoolVarP(&forceRecreate, "force", "f", false, "Force recreate volume")
	CreateVolumeCommand.Flags().StringVarP(&nfsOptions, "opt", "o", "", "Options for nfs")
	CreateVolumeCommand.Flags().StringVarP(&name, "name", "n", "", "Name of volume")
}

var CreateVolumeCommand = cobra.Command{
	Use:   "create-volume",
	Short: "Create docker volume for nfs mount.",
	Run: func(cmd *cobra.Command, args []string) {

		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		cwd := dir

		if name == "" {
			paths := strings.Split(cwd, "/")
			name = strings.Join(paths, "_")[1:]
			name = strings.Replace(name, ".", "_", -1)
			name = strings.Replace(name, "-", "_", -1)
		}

		var execute func() error
		execute = volumeCmd.NewDockerVolumeInspectCommand(name)
		inspectErr := execute()
		if inspectErr != nil {
			if !common.IsVolumeExists(inspectErr) {
				log.Fatal(inspectErr)
			}
			if common.IsVolumeExists(inspectErr) && !forceRecreate {
				log.Print(inspectErr)
				log.Fatalf("failed to create volume %s, use -f option to force create", name)
			} else {
				log.Printf("volume (%s) already exists, force recreate", name)
				execute = volumeCmd.NewDockerVolumeRmCommand(name)
				if err := execute(); err != nil {
					log.Fatal(err)
				}
			}
		}
		execute = volumeCmd.NewDockerVolumeCreateCommand(cwd, name, nfsOptions)
		if err := execute(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success create volume %s in %s", name, cwd)

	},
}
