package cli_acmeaccount_test

import (
	"testing"

	_ "github.com/Telmate/proxmox-api-go/cli/command/commands"
	"github.com/Telmate/proxmox-api-go/proxmox"
	cliTest "github.com/Telmate/proxmox-api-go/test/cli"
)

func Test_AcmeAccount_0_Cleanup(t *testing.T) {
	test := cliTest.Run(t, "-i delete acmeaccount test-0", cliTest.Variables())
	test.ErrorContains("test-0")
}

func Test_AcmeAccount_0_Set(t *testing.T) {
	test := cliTest.Run(t, "-i create acmeaccount test-0", cliTest.Variables(),
		proxmox.ConfigAcmeAccount{
			Contact: []string{
				"a@nonexistantdomain.com",
				"b@nonexistantdomain.com",
				"c@nonexistantdomain.com",
				"d@nonexistantdomain.com",
			},
			Directory: "https://acme-staging-v02.api.letsencrypt.org/directory",
			Tos:       true,
		})
	test.NoError()
	test.Contains("(test-0)")
}

func Test_AcmeAccount_0_Get(t *testing.T) {
	test := cliTest.Run(t, "-i get acmeaccount test-0", cliTest.Variables())
	test.NoError()
	test.JsonEqual(proxmox.ConfigAcmeAccount{
		Name: "test-0",
		Contact: []string{
			"a@nonexistantdomain.com",
			"b@nonexistantdomain.com",
			"c@nonexistantdomain.com",
			"d@nonexistantdomain.com",
		},
		Directory: "https://acme-staging-v02.api.letsencrypt.org/directory",
		Tos:       true,
	})
}

func Test_AcmeAccount_0_Delete(t *testing.T) {
	test := cliTest.Run(t, "-i delete acmeaccount test-0", cliTest.Variables())
	test.NoError()
	test.Contains("test-0")
}
