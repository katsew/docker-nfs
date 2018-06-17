package main

import (
	"os"

	"os/user"

	"log"

	"bufio"

	"os/exec"

	"flag"

	exportsCmd "github.com/katsew/docker-nfs/pkg/cmd/exports"
	nfsConfCmd "github.com/katsew/docker-nfs/pkg/cmd/nfsconf"
	"github.com/katsew/docker-nfs/pkg/common"
	"github.com/katsew/docker-nfs/pkg/exports"
	"github.com/katsew/docker-nfs/pkg/nfsconf"
)

var (
	addr              = flag.String("addr", "localhost", "address for /etc/exports")
	force             = flag.Bool("f", false, "force execute")
	version           = flag.Bool("v", false, "output current version of docker-nfs")
	verboseLevelInfo  = flag.Bool("vv", false, "verbose output level with info")
	verboseLevelDebug = flag.Bool("vvv", false, "verbose output level with debug")
	cwd               string
	Version           = "0.0.0"
)

func init() {

	// Prepare flag default and parse
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cwd = dir
	flag.Parse()

	// Initialize log configuration
	if *verboseLevelDebug {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	} else if *verboseLevelInfo {
		log.SetFlags(log.LstdFlags)
	} else {
		log.SetFlags(0)
	}
	log.SetPrefix("nfsauto | ")

	if *version {
		log.SetPrefix("nfsauto - ")
		log.Printf("version %s", Version)
		os.Exit(0)
	}

}

func main() {

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
		if !*force {
			rd := bufio.NewReader(os.Stdin)
			log.Printf("do you wish to continue? [Y/n]: ")
			ans, _ := rd.ReadString('\n')
			if ans != "Y\n" {
				log.Fatal("answer is not Y, stop process")
			}
		}
	}

	fi, err := os.Stat(exports.PathToExports)
	if err != nil && os.IsExist(err) {
		log.Fatal(err)
	}

	var execute func() error
	var successUpdateExports = false
	if fi != nil {
		execute = exportsCmd.NewAppendCommand(cwd, *addr, uid, gid)
	} else {
		execute = exportsCmd.NewCreateCommand(cwd, *addr, uid, gid)
	}
	if err = execute(); err != nil {
		if common.IsConfigurationExist(err) {
			log.Print(err)
			log.Print("nothing to do")
		} else {
			log.Fatal(err)
		}
	} else {
		successUpdateExports = true
	}

	var successUpdateNfsConf = false
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
	} else {
		successUpdateNfsConf = true
	}

	if successUpdateExports || successUpdateNfsConf {
		log.Print("success update config file, restart nfsd")
		cmd := exec.Command("/bin/sh", "-c", "sudo nfsd restart")
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		log.Print("success restart nfsd")
	}

}
