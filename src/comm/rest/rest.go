package rest

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Suc(data any) gin.H {
	return gin.H{
		"code": 0,
		"data": data,
		"date": time.Now(),
		"msg":  "成功",
	}
}

func Fail(msg any) gin.H {
	return gin.H{
		"code": -1,
		"msg":  msg,
		"data": "",
		"date": time.Now(),
	}
}
