package nfsconf

import (
	"errors"
	"fmt"
	"strings"
)

const (
	PathToNfsConf = "/etc/nfs.conf"
)

type NfsConf []*Configuration

func (n *NfsConf) AlreadyExists(c *Configuration) bool {
	return n.ConfigKeyExists(c.ConfigKey)
}

func (n *NfsConf) ConfigKeyExists(key string) bool {
	for _, conf := range *n {
		if conf.ConfigKey == key {
			return true
		}
	}
	return false
}

func (n *NfsConf) String() string {
	configs := *n
	var str []string
	for i := 0; i < len(configs); i++ {
		str = append(str, fmt.Sprintf("%s\n", configs[i].String()))
	}
	return strings.Join(str, "")
}

type Configuration struct {
	Comment     string
	ConfigKey   string
	ConfigValue string
}

func Parse(line string) (*Configuration, error) {
	if isComment(line) {
		return &Configuration{Comment: line}, nil
	}
	trimmed := strings.Trim(line, " ")
	splitted := strings.Split(trimmed, "=")
	if len(splitted) != 2 {
		return nil, errors.New("length mismatched")
	}
	return &Configuration{
		ConfigKey:   strings.Trim(splitted[0], " "),
		ConfigValue: strings.Trim(splitted[1], " "),
	}, nil
}

func (c *Configuration) String() string {
	if c.Comment != "" {
		return c.Comment
	}
	return fmt.Sprintf("%s = %s", c.ConfigKey, c.ConfigValue)
}

func isComment(line string) bool {
	return strings.Index(line, "#") == 0
}
