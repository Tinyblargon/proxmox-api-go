package content_template_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	_ "github.com/Telmate/proxmox-api-go/cli/command/commands"
	"github.com/Telmate/proxmox-api-go/proxmox"
	cliTest "github.com/Telmate/proxmox-api-go/test/cli"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const storage string = "local"

func checkIfTemplateDoesNotExist(t *testing.T, template, node, storage string) {
	test := cliTest.Run(t, "-i list files "+strings.Join([]string{cliTest.FirstNode, storage, string(proxmox.ContentType_Template)}, " "), cliTest.Variables())
	test.NoError()
	test.NotContains(template)
}

func Test_ContentTemplate_Download_Cleanup(t *testing.T) {
	test := cliTest.Run(t, "-i delete file "+strings.Join([]string{cliTest.FirstNode, storage, string(proxmox.ContentType_Template), cliTest.DownloadedLXCTemplate}, " "), cliTest.Variables())
	test.NoError()
	checkIfTemplateDoesNotExist(t, cliTest.DownloadedLXCTemplate, cliTest.FirstNode, storage)
}

func Test_ContentTemplate_Download(t *testing.T) {
	test := cliTest.Run(t, "-i content template download "+strings.Join([]string{cliTest.FirstNode, storage, cliTest.DownloadedLXCTemplate}, " "), cliTest.Variables())
	test.NoError()
	test.Contains("(" + cliTest.DownloadedLXCTemplate + ")")
}

func Test_ContentTemplate_List(t *testing.T) {
	test := cliTest.Run(t, "-i list files "+strings.Join([]string{cliTest.FirstNode, storage, string(proxmox.ContentType_Template)}, " "), cliTest.Variables())
	test.NoError()
	test.NotEmpty()
	var data []*proxmox.Content_FileProperties
	require.NoError(t, json.Unmarshal([]byte(test.GetOutput()), &data))
	assert.Equal(t, cliTest.DownloadedLXCTemplate, data[0].Name)
	assert.NotEqual(t, "", data[0].Format)
	assert.Greater(t, data[0].Size, uint(0))
	assert.Greater(t, data[0].CreationTime, time.UnixMilli(0))
}

func Test_ContentTemplate_Download_Delete(t *testing.T) {
	test := cliTest.Run(t, "-i delete file "+cliTest.FirstNode+" "+storage+" "+string(proxmox.ContentType_Template)+" "+cliTest.DownloadedLXCTemplate, cliTest.Variables())
	test.NoError()
	test.Contains(cliTest.DownloadedLXCTemplate)
}

func Test_ContentTemplate_Existence_Removed_1(t *testing.T) {
	checkIfTemplateDoesNotExist(t, cliTest.DownloadedLXCTemplate, cliTest.FirstNode, storage)
}
