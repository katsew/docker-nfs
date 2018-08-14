package exports

import (
	"fmt"
	"io/ioutil"

	"github.com/katsew/docker-nfs/pkg/exports"
)

func NewCreateCommand(path string, addr string, uid string, gid string, readOnly bool) func() error {
	return func() error {
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
		err := ioutil.WriteFile(exports.PathToExports, []byte(out.String()), 0644)
		return err
	}
}
