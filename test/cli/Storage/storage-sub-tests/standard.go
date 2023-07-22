package storagesubtests

import (
	"testing"

	_ "github.com/Telmate/proxmox-api-go/cli/command/commands"
	"github.com/Telmate/proxmox-api-go/proxmox"
	cliTest "github.com/Telmate/proxmox-api-go/test/cli"
)

func Cleanup(name string, t *testing.T) {
	test := cliTest.Run(t, "-i delete storage "+name, cliTest.Variables())
	test.ErrorContains(name)
}

func Delete(name string, t *testing.T) {
	test := cliTest.Run(t, "-i delete storage "+name, cliTest.Variables())
	test.NoError()
	test.Contains(name)
}

func Get(s *proxmox.ConfigStorage, name string, t *testing.T) {
	test := cliTest.Run(t, "-i get storage "+name, cliTest.Variables())
	test.NoError()
	test.JsonEqual(s)
}

func Create(s *proxmox.ConfigStorage, name string, t *testing.T) {
	createOrUpdate(s, name, "create", t)
}

func Update(s *proxmox.ConfigStorage, name string, t *testing.T) {
	createOrUpdate(s, name, "update", t)
}

func createOrUpdate(s *proxmox.ConfigStorage, name, command string, t *testing.T) {
	test := cliTest.Run(t, "-i "+command+" storage "+name, cliTest.Variables(), s)
	test.NoError()
	test.Contains("(" + name + ")")
}
