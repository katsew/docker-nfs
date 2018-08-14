// TODO: find more better package to put error struct
package common

import (
	"errors"
	"fmt"
)

var (
	ErrConfigurationExist = errors.New("configuration already exists")
	ErrInvalidLength      = errors.New("length mismatched")
	ErrVolumeAlreadyExist = errors.New("volume already exists")
)

type DockerVolumeAlreadyExists struct {
	Name   string
	Device string
	Option string
	Err    error
}

func (e DockerVolumeAlreadyExists) Error() string {
	return fmt.Sprintf("[%s] (%s) %s with option [%s]", e.Device, e.Name, e.Err.Error(), e.Option)
}

type ConfigurationIsExists struct {
	FilePath string
	Config   string
	Err      error
}

func (e ConfigurationIsExists) Error() string {
	return fmt.Sprintf("[%s] (%s) %s", e.FilePath, e.Config, e.Err.Error())
}

func IsConfigurationExist(err error) bool {
	_, yes := err.(ConfigurationIsExists)
	return yes
}

func IsVolumeExists(err error) bool {
	_, yes := err.(DockerVolumeAlreadyExists)
	return yes
}
