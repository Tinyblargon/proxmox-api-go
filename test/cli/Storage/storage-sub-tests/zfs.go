package storagesubtests

import (
	"testing"

	"github.com/Telmate/proxmox-api-go/proxmox"
)

func ZFSFull() *proxmox.ConfigStorage {
	return &proxmox.ConfigStorage{
		Enable: true,
		Nodes:  []string{"pve"},
		Type:   "zfs",
		ZFS: &proxmox.ConfigStorageZFS{
			Pool:          "test-pool",
			Blocksize:     proxmox.PointerString("4k"),
			Thinprovision: true,
		},
		Content: &proxmox.ConfigStorageContent{
			Container: proxmox.PointerBool(true),
			DiskImage: proxmox.PointerBool(true),
		},
	}
}

func ZFSEmpty() *proxmox.ConfigStorage {
	return &proxmox.ConfigStorage{
		Type: "zfs",
		ZFS: &proxmox.ConfigStorageZFS{
			Pool: "test-pool",
		},
		Content: &proxmox.ConfigStorageContent{
			DiskImage: proxmox.PointerBool(true),
		},
	}
}

func ZFSGetFull(name string, t *testing.T) {
	s := ZFSFull()
	s.ID = name
	Get(s, name, t)
}

func ZFSGetEmpty(name string, t *testing.T) {
	s := ZFSEmpty()
	s.ID = name
	s.ZFS.Blocksize = proxmox.PointerString("8k")
	s.Content.Container = proxmox.PointerBool(false)
	Get(s, name, t)
}
