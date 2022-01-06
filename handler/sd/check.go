package sd

import (
	"github.com/gin-gonic/gin"
	"minepin-backend/pkg/health"
	"net/http"
)

func HealthCheck(c *gin.Context) {
	message := "OK"
	diskMessage,_ := health.DiskStatus()
	cpuMessage,_,_ := health.CpuStatus()
	ramMessage,_ := health.RAMStatus()

	c.String(http.StatusOK, "\n"+
		message+"\n"+
		"CPU:	"+cpuMessage+"\n"+
		"RAM:	"+ramMessage+"\n"+
		"DISK:	"+diskMessage)
}

func DiskCheck(c *gin.Context) {
	message, usedPercent := health.DiskStatus()

	status := http.StatusOK

	if usedPercent >= 95 {
		status = http.StatusInternalServerError
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
	}

	c.String(status, message)
}

func CPUCheck(c *gin.Context) {
	message, cores, l5 := health.CpuStatus()

	status := http.StatusOK

	if l5 >= float64(cores-1) {
		status = http.StatusInternalServerError
	} else if l5 >= float64(cores-2) {
		status = http.StatusTooManyRequests
	}

	c.String(status, "\n"+message)
}

func RAMCheck(c *gin.Context) {
	message, usedPercent := health.RAMStatus()

	status := http.StatusOK

	if usedPercent >= 95 {
		status = http.StatusInternalServerError
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
	}

	c.String(status, "\n"+message)
}
