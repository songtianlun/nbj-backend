package health

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func DiskStatus() (string, int) {
	u, _ := disk.Usage("/")

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	tag := "OK"

	if usedPercent >= 95 {
		tag = "CRITICAL"
	} else if usedPercent >= 90 {
		tag = "WARNING"
	}

	message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", tag, usedMB, usedGB, totalMB, totalGB, usedPercent)
	return message, usedPercent
}

func CpuStatus() (string, int, float64) {
	cores, _ := cpu.Counts(false)

	a, _ := load.Avg()
	l1 := a.Load1
	l5 := a.Load5
	l15 := a.Load15

	tag := "OK"

	if l5 >= float64(cores) {
		tag = "CRITICAL"
	} else if l5 >= float64(cores-1) {
		tag = "WARNING"
	}

	message := fmt.Sprintf("%s - Load average: %.2f, %.2f, %.2f | Cores: %d", tag, l1, l5, l15, cores)
	return message, cores, l5
}

func RAMStatus() (string, int) {
	u, _ := mem.VirtualMemory()

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	text := "OK"

	if usedPercent >= 95 {
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		text = "WARNING"
	}

	message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)

	return message, usedPercent
}
