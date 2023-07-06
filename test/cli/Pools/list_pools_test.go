package cli_pool_test

import (
	"testing"

	cliTest "github.com/Telmate/proxmox-api-go/test/cli"
)

// Test0
func Test_Pool_0_Cleanup(t *testing.T) {
	test := cliTest.Run(t, "-i delete pool test-pool0", cliTest.Variables())
	test.Error()
}

func Test_Pool_0_Create_Without_Comment(t *testing.T) {
	test := cliTest.Run(t, "-i create pool test-pool0", cliTest.Variables())
	test.NoError()
	test.Contains("(test-pool0)")
}

func Test_Pool_0_List(t *testing.T) {
	test := cliTest.Run(t, "-i list pools", cliTest.Variables())
	test.NoError()
	test.Contains(`"test-pool0"`)
}

func Test_Pool_0_Get_Without_Comment(t *testing.T) {
	test := cliTest.Run(t, "-i get pool test-pool0", cliTest.Variables())
	test.NoError()
	test.NotContains(`"comment"`)
}

func Test_Pool_0_Update_Comment(t *testing.T) {
	test := cliTest.Run(t, `-i update poolcomment test-pool0 "this is a comment"`, cliTest.Variables())
	test.NoError()
	test.Contains("(test-pool0)")
}

func Test_Pool_0_Get_With_Comment(t *testing.T) {
	test := cliTest.Run(t, "-i get pool test-pool0", cliTest.Variables())
	test.NoError()
	test.Contains(`"this is a comment"`)
}

func Test_Pool_0_Delete(t *testing.T) {
	test := cliTest.Run(t, "-i delete pool test-pool0", cliTest.Variables())
	test.NoError()
	test.Contains("(test-pool0)")
}

func Test_Pool_0_Removed(t *testing.T) {
	test := cliTest.Run(t, "-i list pools", cliTest.Variables())
	test.NoError()
	test.NotContains(`"test-pool0"`)
}

// Test1
func Test_Pool_1_Cleanup(t *testing.T) {
	test := cliTest.Run(t, "-i delete pool test-pool1", cliTest.Variables())
	test.Error()
}

func Test_Pool_1_Create_With_Comment(t *testing.T) {
	test := cliTest.Run(t, `-i create pool test-pool1 "This is a comment"`, cliTest.Variables())
	test.NoError()
	test.Contains("(test-pool1)")
}

func Test_Pool_1_Get_With_Comment(t *testing.T) {
	test := cliTest.Run(t, "-i get pool test-pool1", cliTest.Variables())
	test.NoError()
	test.Contains(`"This is a comment"`)
}

func Test_Pool_1_Update_Comment(t *testing.T) {
	test := cliTest.Run(t, "-i update poolcomment test-pool1", cliTest.Variables())
	test.NoError()
	test.Contains("(test-pool1)")
}

func Test_Pool_1_Get_Without_Comment(t *testing.T) {
	test := cliTest.Run(t, "-i get pool test-pool1", cliTest.Variables())
	test.NoError()
	test.NotContains(`"comment"`)
}

func Test_Pool_1_Delete(t *testing.T) {
	test := cliTest.Run(t, "-i delete pool test-pool1", cliTest.Variables())
	test.NoError()
	test.Contains("(test-pool1)")
}

func Test_Pool_1_Removed(t *testing.T) {
	test := cliTest.Run(t, "-i list pools", cliTest.Variables())
	test.NoError()
	test.NotContains(`"test-pool1"`)
}
