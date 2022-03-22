package utils

import (
	"github.com/gin-gonic/gin"
	"mingin/pkg/errno"
	"os"
	"strconv"
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

func GetUint64ByContext(c *gin.Context, k string) (uint64, error) {
	if pid := c.Param(k); pid != "" {
		n, err := strconv.ParseUint(pid, 10, 64)
		if err != nil {
			return 0, err
		} else {
			return n, nil
		}
	}
	return 0, errno.ErrParamKey
}
