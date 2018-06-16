package nfsconf

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/katsew/docker-nfs/pkg/common"
	"github.com/katsew/docker-nfs/pkg/nfsconf"
)

func NewAppendCommand() func() error {
	return func() error {
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
				return err
			}

			conf, err := nfsconf.Parse(string(l))
			if err != nil {
				return err
			}
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
			return common.ConfigurationIsExists{
				FilePath: "/etc/nfs.conf",
				Config:   c.String(),
				Err:      common.ErrConfigurationExist,
			}
		}
		_, err = f.WriteString(out.String())
		if err != nil {
			return err
		}
		return nil
	}
}
