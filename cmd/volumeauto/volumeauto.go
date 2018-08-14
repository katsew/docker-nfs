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
	name       *string
	version    = flag.Bool("v", false, "output current version of volumeauto")
	Version    = "0.0.0"
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
	name = flag.String("name", fmt.Sprintf("%s_%s", "nfs_auto", projectRoot), "docker volume name")
	flag.Parse()
	log.SetPrefix("volumeauto | ")

	if *version {
		log.SetPrefix("nfsauto - ")
		log.Printf("version %s", Version)
		os.Exit(0)
	}

}

func main() {

	log.Print("[warning] This command has been deprecated. use docker-nfs instead.")
	log.Print("[hint] go get github.com/katsew/docker-nfs")

	var execute func() error
	execute = volumeCmd.NewDockerVolumeRmCommand(*name)
	if err := execute(); err != nil {
		if !*force {
			log.Fatalf("failed to rm volume %s, use -f option to force create", *name)
		}
		log.Printf("failed to rm existing volume %s, continue create one", *name)
	}

	execute = volumeCmd.NewDockerVolumeCreateCommand(cwd, *name, *nfsOptions)
	if err := execute(); err != nil {
		log.Fatal(err)
	}
	log.Printf("success create volume %s in %s", *name, cwd)

}
