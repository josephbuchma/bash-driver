// Package bash implements the Driver interface.
package bash

import (
	"os/exec"

	"github.com/db-journey/migrate/driver"
	"github.com/db-journey/migrate/file"
)

type Driver struct {
}

// make sure our driver still implements the driver.Driver interface
var _ driver.Driver = (*Driver)(nil)

func (driver *Driver) Initialize(url string) error {
	return nil
}

func (driver *Driver) Close() error {
	return nil
}

func (driver *Driver) FilenameExtension() string {
	return "sh"
}

func (driver *Driver) Migrate(f file.File, pipe chan interface{}) {
	defer close(pipe)
	pipe <- f
	return
}

// Version returns the current migration version.
func (driver *Driver) Version() (file.Version, error) {
	return file.Version(0), nil
}

// Versions returns the list of applied migrations.
func (driver *Driver) Versions() (file.Versions, error) {
	return file.Versions{0}, nil
}

// Execute shell script
func (driver *Driver) Execute(commands string) error {
	return exec.Command("sh", "-c", commands).Run()
}

func init() {
	driver.RegisterDriver("bash", &Driver{})
}
