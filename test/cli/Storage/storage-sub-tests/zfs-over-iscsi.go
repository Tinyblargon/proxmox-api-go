package storagesubtests

import (
	"github.com/Telmate/proxmox-api-go/proxmox"
)

func ZFSoverISCSIFull() *proxmox.ConfigStorage {
	return &proxmox.ConfigStorage{
		Enable: true,
		Nodes:  []string{"pve"},
		Type:   "zfs-over-iscsi",
		ZFSoverISCSI: &proxmox.ConfigStorageZFSoverISCSI{
			Portal:        "test-portal",
			Pool:          "test-pool",
			Blocksize:     proxmox.PointerString("8k"),
			Target:        "test-target",
			Thinprovision: true,
			ISCSIprovider: "iet",
		},
		Content: &proxmox.ConfigStorageContent{
			DiskImage: proxmox.PointerBool(true),
		},
	}
}

func ZFSoverISCSIEmpty() *proxmox.ConfigStorage {
	return &proxmox.ConfigStorage{
		Type: "zfs-over-iscsi",
		ZFSoverISCSI: &proxmox.ConfigStorageZFSoverISCSI{
			Portal:        "test-portal",
			Pool:          "test-pool",
			Target:        "test-target",
			ISCSIprovider: "iet",
		},
	}
}
