package nfsconf

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/katsew/docker-nfs/pkg/nfsconf"
)

func NewAppendCommand() func() {
	return func() {
		f, err := os.OpenFile(nfsconf.PathToNfsConf, os.O_RDWR, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		rd := bufio.NewReader(f)
		var exists nfsconf.NfsConf
		for {
			l, _, err := rd.ReadLine()

			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}

			conf, err := nfsconf.Parse(string(l))
			exists = append(exists, conf)
		}
		begin := &nfsconf.Configuration{Comment: fmt.Sprintf("## BEGIN - docker-nfs")}
		c := &nfsconf.Configuration{
			ConfigKey:   "nfs.server.mount.require_resv_port",
			ConfigValue: "0",
		}
		end := &nfsconf.Configuration{Comment: fmt.Sprintf("## END - docker-nfs")}
		out := nfsconf.NfsConf{begin, c, end}
		if exists.AlreadyExists(c) {
			fmt.Printf("\nThis configuration already exists on current nfs.conf.")
			return
		}
		_, err = f.WriteString(out.String())
		if err != nil {
			panic(err)
		}
	}
}
