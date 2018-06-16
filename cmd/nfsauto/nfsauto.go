package main

import (
	"os"

	"os/user"

	"log"

	"bufio"

	exportsCmd "github.com/katsew/docker-nfs/pkg/cmd/exports"
	nfsConfCmd "github.com/katsew/docker-nfs/pkg/cmd/nfsconf"
	"github.com/katsew/docker-nfs/pkg/common"
	"github.com/katsew/docker-nfs/pkg/exports"
	"github.com/katsew/docker-nfs/pkg/nfsconf"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	uid := os.Getenv("SUDO_UID")
	gid := os.Getenv("SUDO_GID")
	if uid == "" || gid == "" {
		log.Print("uid:gid for sudo does not exists, forget `sudo`?")
		u, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		uid = u.Uid
		gid = u.Gid
		log.Printf("uid:gid set to %s:%s, this may break your config!", uid, gid)
		rd := bufio.NewReader(os.Stdin)
		log.Printf("do you wish to continue? [Y/n]: ")
		ans, _ := rd.ReadString('\n')
		if ans != "Y\n" {
			log.Fatal("answer is not Y, stop process")
		}
	}
	var address = "localhost"
	if len(os.Args) > 1 {
		address = os.Args[1]
	}
	fi, err := os.Stat(exports.PathToExports)
	if err != nil && os.IsExist(err) {
		log.Fatal(err)
	}

	var execute func() error
	if fi != nil {
		execute = exportsCmd.NewAppendCommand(dir, address, uid, gid)
	} else {
		execute = exportsCmd.NewCreateCommand(dir, address, uid, gid)
	}
	if err = execute(); err != nil {
		if common.IsConfigurationExist(err) {
			log.Print(err)
			log.Print("nothing to do")
		} else {
			log.Fatal(err)
		}
	}

	fi, err = os.Stat(nfsconf.PathToNfsConf)
	if err != nil && os.IsExist(err) {
		log.Fatal(err)
	}
	if fi != nil {
		execute = nfsConfCmd.NewAppendCommand()
	} else {
		execute = nfsConfCmd.NewCreateCommand()
	}
	if err = execute(); err != nil {
		if common.IsConfigurationExist(err) {
			log.Print(err)
			log.Print("nothing to do")
		} else {
			log.Fatal(err)
		}
	}
}
