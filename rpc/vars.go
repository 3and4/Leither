// vars.go
package rpc

import (
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

const (
	ApiVarVersion       = "ver"           //应用版本
	ApiVarCPU           = "cpu"           //获取cpu的状态
	ApiVarDisk          = "disk"          //获取cpu的状态
	ApiVarVirtualMemory = "virtualmemory" //获取虚拟内存的状态

)

type StatCPU = []float64

// Memory usage statistics. Total, Available and Used contain numbers of bytes
// for human consumption.
//
// The other fields in this struct contain kernel specific values.
// type VirtualMemoryStat struct {
// 	// Total amount of RAM on this system
// 	Total uint64 `json:"total"`

// 	// RAM available for programs to allocate
// 	//
// 	// This value is computed from the kernel specific values.
// 	Available uint64 `json:"available"`

// 	// RAM used by programs
// 	//
// 	// This value is computed from the kernel specific values.
// 	Used uint64 `json:"used"`

// 	// Percentage of RAM used by programs
// 	//
// 	// This value is computed from the kernel specific values.
// 	UsedPercent float64 `json:"usedPercent"`

// 	// This is the kernel's notion of free memory; RAM chips whose bits nobody
// 	// cares about the value of right now. For a human consumable number,
// 	// Available is what you really want.
// 	Free uint64 `json:"free"`

// 	// OS X / BSD specific numbers:
// 	// http://www.macyourself.com/2010/02/17/what-is-free-wired-active-and-inactive-system-memory-ram/
// 	Active   uint64 `json:"active"`
// 	Inactive uint64 `json:"inactive"`
// 	Wired    uint64 `json:"wired"`

// 	// Linux specific numbers
// 	// https://www.centos.org/docs/5/html/5.1/Deployment_Guide/s2-proc-meminfo.html
// 	Buffers uint64 `json:"buffers"`
// 	Cached  uint64 `json:"cached"`
// }

type StatVirtualMemory = mem.VirtualMemoryStat

//	type UsageStat struct {
//		Path              string  `json:"path"`
//		Fstype            string  `json:"fstype"`
//		Total             uint64  `json:"total"`
//		Free              uint64  `json:"free"`
//		Used              uint64  `json:"used"`
//		UsedPercent       float64 `json:"usedPercent"`
//		InodesTotal       uint64  `json:"inodesTotal"`
//		InodesUsed        uint64  `json:"inodesUsed"`
//		InodesFree        uint64  `json:"inodesFree"`
//		InodesUsedPercent float64 `json:"inodesUsedPercent"`
//	}
type StatUsage = disk.UsageStat
