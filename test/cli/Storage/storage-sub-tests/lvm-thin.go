package storagesubtests

import (
	"testing"

	"github.com/Telmate/proxmox-api-go/proxmox"
)

func LVMThinFull() *proxmox.ConfigStorage {
	return &proxmox.ConfigStorage{
		Enable: true,
		Nodes:  []string{"pve"},
		Type:   "lvm-thin",
		LVMThin: &proxmox.ConfigStorageLVMThin{
			VGname:   "pve",
			Thinpool: "data",
		},
		Content: &proxmox.ConfigStorageContent{
			Container: proxmox.PointerBool(true),
			DiskImage: proxmox.PointerBool(true),
		},
	}
}

func LVMThinEmpty() *proxmox.ConfigStorage {
	return &proxmox.ConfigStorage{
		Type: "lvm-thin",
		LVMThin: &proxmox.ConfigStorageLVMThin{
			VGname:   "pve",
			Thinpool: "data",
		},
		Content: &proxmox.ConfigStorageContent{
			Container: proxmox.PointerBool(true),
		},
	}
}

func LVMThinGetFull(name string, t *testing.T) {
	s := LVMThinFull()
	s.ID = name
	Get(s, name, t)
}

func LVMThinGetEmpty(name string, t *testing.T) {
	s := LVMThinEmpty()
	s.ID = name
	s.Content.DiskImage = proxmox.PointerBool(false)
	Get(s, name, t)
}
