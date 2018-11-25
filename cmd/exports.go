package cmd

import (
	"os"

	"os/user"

	"log"

	"bufio"

	"os/exec"

	exportsCmd "github.com/katsew/docker-nfs/pkg/cmd/exports"
	nfsConfCmd "github.com/katsew/docker-nfs/pkg/cmd/nfsconf"
	"github.com/katsew/docker-nfs/pkg/common"
	"github.com/katsew/docker-nfs/pkg/exports"
	"github.com/katsew/docker-nfs/pkg/nfsconf"
	"github.com/spf13/cobra"
)

var (
	addr        string
	forceExport bool
	readOnly    bool
)

func init() {
	ExportsCommand.Flags().BoolVarP(&forceExport, "force", "f", false, "Export address without confirmation")
	ExportsCommand.Flags().BoolVarP(&readOnly, "read-only", "", false, "Export as read only mount")
	ExportsCommand.Flags().StringVarP(&addr, "addr", "a", "localhost", "Address to export")
}

var ExportsCommand = cobra.Command{
	Use:   "exports",
	Short: "Set nfs configuration into /etc/nfs.conf, /etc/exports and restarts nfs.",
	Run: func(cmd *cobra.Command, args []string) {

		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		cwd := dir

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
			if !forceExport {
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
			execute = exportsCmd.NewAppendCommand(cwd, addr, uid, gid, readOnly)
		} else {
			execute = exportsCmd.NewCreateCommand(cwd, addr, uid, gid, readOnly)
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
	},
}
