package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/katsew/docker-nfs/pkg/exports"
)

func NewAppendCommand(path string, addr string, uid string, gid string) func() {
	return func() {
		f, err := os.OpenFile(exports.PathToExports, os.O_RDWR, 0644)
		if err != nil {
			panic(err)
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
				panic(err)
			}

			conf, err := exports.Parse(string(l))
			exists = append(exists, conf)
		}
		begin := &exports.Configuration{Comment: fmt.Sprintf("# BEGIN - docker-nfs %s:%s", uid, gid)}
		c := &exports.Configuration{
			Path: path,
			Addr: addr,
			Uid:  uid,
			Gid:  gid,
		}
		end := &exports.Configuration{Comment: fmt.Sprintf("# END - docker-nfs %s:%s", uid, gid)}
		out := exports.Exports{begin, c, end}
		if exists.AlreadyExists(c) {
			fmt.Printf("\nThis configuration already exists on current exports.")
			return
		}
		_, err = f.WriteString(out.String())
		if err != nil {
			panic(err)
		}
	}
}
