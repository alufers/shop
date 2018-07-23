package shop

import (
	"math/bits"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUintParam(c *gin.Context, name string) uint {
	if okParam, err := strconv.ParseUint(c.Param(name), 10, bits.UintSize); err == nil {
		return uint(okParam)
	}
	return 0
}
