package publicapi

import "github.com/gin-gonic/gin"

// 创建一个新的图书服务
func NewBooksService(e *gin.Engine) {
	{
		e.GET("/books", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"data": []any{"hello", "word", "你好", "世界", 7889},
			})
		})
	}
}
