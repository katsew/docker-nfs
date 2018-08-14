package main

import (
	"github.com/katsew/docker-nfs/cmd"
	"github.com/spf13/cobra"
)

func main() {

	var dockerNfsCommand = &cobra.Command{
		Use: "docker-nfs",
	}

	dockerNfsCommand.AddCommand(&cmd.ExportsCommand)
	dockerNfsCommand.AddCommand(&cmd.CreateVolumeCommand)

	if err := dockerNfsCommand.Execute(); err != nil {
		panic(err)
	}
}
