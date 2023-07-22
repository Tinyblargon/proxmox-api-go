package storagesubtests

import (
	"testing"

	"github.com/Telmate/proxmox-api-go/proxmox"
)

func IscsiFull() *proxmox.ConfigStorage {
	return &proxmox.ConfigStorage{
		Enable: true,
		Nodes:  []string{"pve"},
		Type:   "iscsi",
		ISCSI: &proxmox.ConfigStorageISCSI{
			Portal: "10.20.1.1",
			Target: "target-volume",
		},
		Content: &proxmox.ConfigStorageContent{
			DiskImage: proxmox.PointerBool(true),
		},
	}
}

func IscsiEmpty() *proxmox.ConfigStorage {
	return &proxmox.ConfigStorage{
		Type: "iscsi",
		ISCSI: &proxmox.ConfigStorageISCSI{
			Portal: "10.20.1.1",
			Target: "target-volume",
		},
		Content: &proxmox.ConfigStorageContent{},
	}
}

func IscsiGetFull(name string, t *testing.T) {
	s := IscsiFull()
	s.ID = name
	Get(s, name, t)
}

func IscsiGetEmpty(name string, t *testing.T) {
	s := IscsiEmpty()
	s.ID = name
	s.Content.DiskImage = proxmox.PointerBool(false)
	Get(s, name, t)
}
