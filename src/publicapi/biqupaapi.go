package publicapi

import (
	"BookAnalysisTool/src/comm/rest"
	"BookAnalysisTool/src/service/comm"
	"BookAnalysisTool/src/service/crawler"

	"github.com/gin-gonic/gin"
)

var biqu = crawler.NewBiqupaManagement()

//笔趣阁爬虫api
func NewPiQuPaApi(e *gin.Engine) {
	rg := e.Group("/biqu")
	{
		// 搜索某本书
		rg.GET("/search/book/:name", func(ctx *gin.Context) {
			name := ctx.Param("name")
			book := biqu.BiqupaBookUrls(name)
			ctx.JSON(200, rest.Suc(book))
		})
		// 爬取指定书籍的所有章节信息
		rg.GET("/search/bookforchapter/:url", func(ctx *gin.Context) {
			url := ctx.Param("url")
			c := biqu.BiqupaChapterUrls(url)
			ctx.JSON(200, rest.Suc(c))
		})
		// 爬取指定章节内容
		rg.GET("/serach/chapter/:url", func(ctx *gin.Context) {
			url := ctx.Param("url")
			var chapter comm.Chapter
			chapter.BiQuUrl = url
			biqu.BiqupaOneChapterContent(&chapter)
			ctx.JSON(200, rest.Suc(chapter))
		})
	}
}
