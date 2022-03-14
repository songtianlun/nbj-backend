package utils

import (
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetAddrFromContext(c *gin.Context) (addr string) {
	return strings.Split(c.Request.RemoteAddr, ":")[0]
}
