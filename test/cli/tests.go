package test

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/Telmate/proxmox-api-go/cli"
	_ "github.com/Telmate/proxmox-api-go/cli/command/commands"
	"github.com/stretchr/testify/require"
)

func Login(UserID string, Password string, OTP string, HttpHeaders string) error {
	_, err := cli.Client(ApiUrl, UserID, Password, OTP, HttpHeaders)
	return err
}

// Run the cli command.
// commandInput is optional and only the first instance wil be evaluated.
func Run(t *testing.T, command []string, variables []EnvironmentVariable, commandInput ...any) TestOutput {
	for _, e := range variables {
		os.Setenv(e.Name, e.Value)
	}
	cli.RootCmd.SetArgs(command)
	buffer := new(bytes.Buffer)
	cli.RootCmd.SetOut(buffer)
	setCommandInput(t, commandInput)
	err := cli.RootCmd.Execute()
	out, _ := io.ReadAll(buffer)
	return TestOutput{err: err, output: string(out), t: t}
}

// set the stdin of the command.
func setCommandInput(t *testing.T, jsonIn []any) {
	if len(jsonIn) == 0 {
		return
	}
	switch InputJson := jsonIn[0].(type) {
	case string:
		if InputJson != "" {
			cli.RootCmd.SetIn(strings.NewReader(InputJson))
		}
	default:
		if InputJson != nil {
			tmpJson, err := json.Marshal(InputJson)
			require.NoError(t, err)
			cli.RootCmd.SetIn(strings.NewReader(string(tmpJson)))
		}
	}
}

// override is optional.
func Variables(override ...map[string]string) []EnvironmentVariable {
	defaults := map[string]string{
		ApiUrl_Var:       ApiUrl,
		RootUser_Var:     RootUser,
		RootPassword_Var: RootPassword,
	}
	if len(override) != 0 {
		for key, value := range override[0] {
			defaults[key] = value
		}
	}
	return []EnvironmentVariable{
		{Name: ApiUrl_Var, Value: defaults[ApiUrl_Var]},
		{Name: RootUser_Var, Value: defaults[RootUser_Var]},
		{Name: RootPassword_Var, Value: defaults[RootPassword_Var]},
	}
}

type EnvironmentVariable struct {
	Name  string
	Value string
}

type TestOutput struct {
	output string
	err    error
	t      *testing.T
}

// Returns a copy of the command error as to not accidentally mutate it.
func (o TestOutput) GetErr() error {
	return o.err
}

// Returns a copy of the command output as to not accidentally mutate it.
func (o TestOutput) GetOutput() string {
	return o.output
}
