package proxmox

import (
	"errors"
)

type CpuFlags struct {
	AES        *bool `json:"aes,omitempty"`        // Activate AES instruction set for HW acceleration.
	AmdNoSSB   *bool `json:"amdnossb,omitempty"`   // Notifies guest OS that host is not vulnerable for Spectre on AMD CPUs.
	AmdSSBD    *bool `json:"amdssbd,omitempty"`    // Improves Spectre mitigation performance with AMD CPUs, best used with "VirtSSBD".
	HvEvmcs    *bool `json:"hvevmcs,omitempty"`    // Improve performance for nested virtualization. Only supported on Intel CPUs.
	HvTlbFlush *bool `json:"hvtlbflush,omitempty"` // Improve performance in overcommitted Windows guests. May lead to guest bluescreens on old CPUs.
	Ibpb       *bool `json:"ibpb,omitempty"`       // Allows improved Spectre mitigation with AMD CPUs.
	MdClear    *bool `json:"mdclear,omitempty"`    // Required to let the guest OS know if MDS is mitigated correctly.
	PCID       *bool `json:"pcid,omitempty"`       // Meltdown fix cost reduction on Westmere, Sandy-, and IvyBridge Intel CPUs.
	Pdpe1GB    *bool `json:"pdpe1gb,omitempty"`    // Allow guest OS to use 1GB size pages, if host HW supports it.
	SpecCtrl   *bool `json:"specctrl,omitempty"`   // Allows improved Spectre mitigation with Intel CPUs.
	SSBD       *bool `json:"ssbd,omitempty"`       // Protection for "Speculative Store Bypass" for Intel models.
	VirtSSBD   *bool `json:"cirtssbd,omitempty"`   // Basis for "Speculative Store Bypass" protection for AMD models.
}

// min value 0 is unlimited, max value of 128
type CpuLimit uint8

const (
	CpuLimit_Error_UpperBound string = "maximum value of CpuLimit is 128"
)

func (limit CpuLimit) Validate() error {
	if limit > 128 {
		return errors.New(CpuLimit_Error_UpperBound)
	}
	return nil
}

// min value 1, max value 4
type CpuSockets uint8

const (
	CpuSockets_Error_LowerBound string = "minimum value of CpuSockets is 1"
	CpuSockets_Error_UpperBound string = "maximum value of CpuSockets is 4"
)

func (sockets CpuSockets) Validate() error {
	if sockets < 1 {
		return errors.New(CpuSockets_Error_LowerBound)
	}
	if sockets > 4 {
		return errors.New(CpuSockets_Error_UpperBound)
	}
	return nil
}

// enum
type CpuType string

// min value 0 is unset, max value of 262144
type CpuUnits uint32

// min value 0 is unset, max value 512. is QemuCpuCores * CpuSockets
type CpuVirtualCores uint16

type QemuCPU struct {
	Affinity     []uint          `json:"affinity,omitempty"`
	Cores        QemuCpuCores    `json:"cores"`
	Flags        CpuFlags        `json:"flags"`
	Limit        CpuLimit        `json:"limit,omitempty"`
	Numa         bool            `json:"numa"`
	Sockets      CpuSockets      `json:"sockets"`
	Type         CpuType         `json:"type,omitempty"`
	Units        CpuUnits        `json:"units,omitempty"`
	VirtualCores CpuVirtualCores `json:"vcores,omitempty"`
}

// min value 1, max value of 128
type QemuCpuCores uint8
