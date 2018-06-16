package main

import (
	"os"

	"os/user"

	exportsCmd "github.com/katsew/docker-nfs/pkg/cmd/exports"
	nfsConfCmd "github.com/katsew/docker-nfs/pkg/cmd/nfsconf"
	"github.com/katsew/docker-nfs/pkg/exports"
	"github.com/katsew/docker-nfs/pkg/nfsconf"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	uid := os.Getenv("SUDO_UID")
	gid := os.Getenv("SUDO_GID")
	if uid == "" || gid == "" {
		u, err := user.Current()
		if err != nil {
			panic(err)
		}
		uid = u.Uid
		gid = u.Gid
	}
	var address = "localhost"
	if len(os.Args) > 1 {
		address = os.Args[1]
	}

	fi, err := os.Stat(exports.PathToExports)
	if err != nil && os.IsExist(err) {
		panic(err)
	}
	var execute func()
	if fi != nil {
		execute = exportsCmd.NewAppendCommand(dir, address, uid, gid)
	} else {
		execute = exportsCmd.NewCreateCommand(dir, address, uid, gid)
	}
	execute()

	fi, err = os.Stat(nfsconf.PathToNfsConf)
	if err != nil && os.IsExist(err) {
		panic(err)
	}
	if fi != nil {
		execute = nfsConfCmd.NewAppendCommand()
	} else {
		execute = nfsConfCmd.NewCreateCommand()
	}
	execute()
}
