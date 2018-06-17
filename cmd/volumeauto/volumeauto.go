package main

import (
	"os"

	"log"

	"flag"

	"fmt"
	"strings"

	volumeCmd "github.com/katsew/docker-nfs/pkg/cmd/docker/volume"
)

var (
	force      = flag.Bool("f", false, "force execute")
	nfsOptions = flag.String("o", "", "nfs options for docker volume")
	cwd        string
	volumeName *string
)

func init() {

	// Prepare flag default and parse
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cwd = dir
	splits := strings.Split(dir, "/")
	projectRoot := splits[len(splits)-1]
	volumeName = flag.String("volume", fmt.Sprintf("%s_%s", "nfs_auto", projectRoot), "docker volume name")
	flag.Parse()
	log.SetPrefix("volumeauto | ")

}

func main() {

	var execute func() error
	execute = volumeCmd.NewDockerVolumeRmCommand(*volumeName)
	if err := execute(); err != nil {
		if !*force {
			log.Fatalf("failed to rm volume %s, use -f option to force create", *volumeName)
		}
		log.Printf("failed to rm existing volume %s, continue create one", *volumeName)
	}

	execute = volumeCmd.NewDockerVolumeCreateCommand(cwd, *volumeName, *nfsOptions)
	if err := execute(); err != nil {
		log.Fatal(err)
	}
	log.Printf("success create volume %s in %s", *volumeName, cwd)

}
