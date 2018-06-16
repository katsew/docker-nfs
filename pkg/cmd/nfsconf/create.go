package nfsconf

import (
	"fmt"
	"io/ioutil"

	"log"

	"github.com/katsew/docker-nfs/pkg/nfsconf"
)

func NewCreateCommand() func() error {
	return func() error {
		begin := &nfsconf.Configuration{Comment: fmt.Sprintf("## BEGIN - docker-nfs")}
		c := &nfsconf.Configuration{
			ConfigKey:   "nfs.server.mount.require_resv_port",
			ConfigValue: "0",
		}
		end := &nfsconf.Configuration{Comment: fmt.Sprintf("## END - docker-nfs")}
		out := nfsconf.NfsConf{begin, c, end}
		log.Print("write file")
		err := ioutil.WriteFile(nfsconf.PathToNfsConf, []byte(out.String()), 0644)
		return err
	}
}
