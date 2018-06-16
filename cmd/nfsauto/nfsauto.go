package main

import (
	"os"

	"os/user"

	"github.com/katsew/docker-nfs/pkg/cmd"
	"github.com/katsew/docker-nfs/pkg/exports"
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
		execute = cmd.NewAppendCommand(dir, address, uid, gid)
	} else {
		execute = cmd.NewCreateCommand(dir, address, uid, gid)
	}
	execute()
}
