package cli_acmeaccount_test

import (
	"testing"

	_ "github.com/Telmate/proxmox-api-go/cli/command/commands"
	cliTest "github.com/Telmate/proxmox-api-go/test/cli"
)

func Test_AcmeAccount_1_Cleanup(t *testing.T) {
	test := cliTest.Run(t, "-i delete acmeaccount test-1", cliTest.Variables())
	test.ErrorContains("test-1")
}

func Test_AcmeAccount_1_Set(t *testing.T) {
	test := cliTest.Run(t, "-i create acmeaccount test-1", cliTest.Variables(), `
	{
		"contact": [
			"a@nonexistantdomain.com"
		],
		"directory": "https://acme-staging-v02.api.letsencrypt.org/directory",
		"tos": true
	}`)
	test.NoError()
	test.Contains("(test-1)")
}

func Test_AcmeAccount_1_Get(t *testing.T) {
	test := cliTest.Run(t, "-i get acmeaccount test-1", cliTest.Variables())
	test.NoError()
	test.JsonEqual(`
	{
		"name": "test-1",
		"contact": [
			"a@nonexistantdomain.com"
		],
		"directory": "https://acme-staging-v02.api.letsencrypt.org/directory",
		"tos": true
	}`)
}

func Test_AcmeAccount_1_Delete(t *testing.T) {
	test := cliTest.Run(t, "-i delete acmeaccount test-1", cliTest.Variables())
	test.NoError()
	test.Contains("test-1")
}
