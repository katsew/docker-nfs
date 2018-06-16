package exports

import (
	"fmt"
	"strings"

	"github.com/katsew/docker-nfs/pkg/common"
)

const (
	PathToExports = "/etc/exports"
)

type Exports []*Configuration

func (e *Exports) String() string {
	configs := *e
	var str []string
	for i := 0; i < len(configs); i++ {
		str = append(str, fmt.Sprintf("%s\n", configs[i].String()))
	}
	return strings.Join(str, "")
}

func (e *Exports) AddrExists(addr string) bool {
	for _, eps := range *e {
		if addr == eps.Addr {
			return true
		}
	}
	return false
}

func (e *Exports) PathExists(path string) bool {
	for _, eps := range *e {
		if path == eps.Path {
			return true
		}
	}
	return false
}

func (e *Exports) UidExists(uid string) bool {
	for _, eps := range *e {
		if uid == eps.Uid {
			return true
		}
	}
	return false
}

func (e *Exports) GidExists(gid string) bool {
	for _, eps := range *e {
		if gid == eps.Gid {
			return true
		}
	}
	return false
}

func (e *Exports) AlreadyExists(c *Configuration) bool {
	return e.PathExists(c.Path) &&
		e.AddrExists(c.Addr) &&
		e.UidExists(c.Uid) &&
		e.GidExists(c.Gid)
}

type Configuration struct {
	Path      string
	Addr      string
	Uid       string
	Gid       string
	OptAllDir string
	Comment   string
}

func (c *Configuration) String() string {
	if c.Comment != "" {
		return c.Comment
	}
	return fmt.Sprintf("\"%s\" %s -alldirs -mapall=%s:%s", c.Path, c.Addr, c.Uid, c.Gid)
}

func Parse(line string) (*Configuration, error) {
	if isComment(line) {
		return &Configuration{Comment: line}, nil
	}
	elements := strings.Split(line, " ")
	var conf = &Configuration{}
	for i, e := range elements {
		trimmed := strings.Trim(e, " ")
		if i == 0 {
			unwrapped := unwrapQuotes(trimmed)
			conf.Path = unwrapped
		}
		if i == 1 {
			conf.Addr = trimmed
		}
		if isAllDirsOption(trimmed) {
			conf.OptAllDir = trimmed
		}
		if isMapAllOption(trimmed) {
			trimmed = strings.TrimPrefix(trimmed, "-mapall=")
			splits := strings.Split(trimmed, ":")
			if len(splits) != 2 {
				return nil, common.ErrInvalidLength
			}
			conf.Uid = splits[0]
			conf.Gid = splits[1]
		}
	}
	return conf, nil
}

func isComment(line string) bool {
	return strings.Index(line, "#") == 0
}
func isAllDirsOption(elem string) bool {
	return strings.HasPrefix(elem, "-alldirs")
}

func isMapAllOption(elem string) bool {
	return strings.HasPrefix(elem, "-mapall=")
}

func unwrapQuotes(elem string) string {
	return strings.Replace(elem, "\"", "", -1)
}
