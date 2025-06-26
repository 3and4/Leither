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
type StatUsage = disk.UsageStat

type StatVirtualMemory = mem.VirtualMemoryStat

// type UsageStat struct {
// 	Path              string  `json:"path"`
// 	Fstype            string  `json:"fstype"`
// 	Total             uint64  `json:"total"`
// 	Free              uint64  `json:"free"`
// 	Used              uint64  `json:"used"`
// 	UsedPercent       float64 `json:"usedPercent"`
// 	InodesTotal       uint64  `json:"inodesTotal"`
// 	InodesUsed        uint64  `json:"inodesUsed"`
// 	InodesFree        uint64  `json:"inodesFree"`
// 	InodesUsedPercent float64 `json:"inodesUsedPercent"`
// }
