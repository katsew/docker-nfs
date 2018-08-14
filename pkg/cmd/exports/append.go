package exports

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/katsew/docker-nfs/pkg/common"
	"github.com/katsew/docker-nfs/pkg/exports"
)

func NewAppendCommand(path string, addr string, uid string, gid string, readOnly bool) func() error {
	return func() error {

		f, err := os.OpenFile(exports.PathToExports, os.O_RDWR, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		rd := bufio.NewReader(f)
		var exists exports.Exports
		for {
			l, _, err := rd.ReadLine()

			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}

			conf, err := exports.Parse(string(l))
			exists = append(exists, conf)
		}
		begin := &exports.Configuration{Comment: fmt.Sprintf("# BEGIN - docker-nfs %s:%s", uid, gid)}
		c := &exports.Configuration{
			Path:     path,
			Addr:     addr,
			Uid:      uid,
			Gid:      gid,
			ReadOnly: readOnly,
		}
		end := &exports.Configuration{Comment: fmt.Sprintf("# END - docker-nfs %s:%s", uid, gid)}
		out := exports.Exports{begin, c, end}
		if exists.AlreadyExists(c) {
			return common.ConfigurationIsExists{
				FilePath: exports.PathToExports,
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
