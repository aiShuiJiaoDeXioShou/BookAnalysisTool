package publicapi

import (
	"BookAnalysisTool/src/parse/crawler"

	"github.com/gin-gonic/gin"
)

// 创建一个新的爬取起点中文网服务
func NewQiDianSort(e *gin.Engine) {
	{
		qd := e.Group("/qidian")
		qd.GET("/paihan/:name", paihan)
	}
}

// 获取月票榜的数据
func paihan(ctx *gin.Context) {
	var name = ctx.Param("name")
	var qidiantype string
	switch name {
	case "月票榜":
		qidiantype = crawler.NOVEL_RANKING
	case "畅销榜":
		qidiantype = crawler.NOVEL_SALES_RANKING
	case "阅读指数":
		qidiantype = crawler.NOVEL_READ_RANKING
	case "推荐榜":
		qidiantype = crawler.NOVEL_RECOMMEND_RANKING
	case "收藏榜":
		qidiantype = crawler.NOVEL_COLLECT_RANKING
	case "签约作者榜":
		qidiantype = crawler.NOVEL_NEW_AUTHOR_RANKING
	case "公共作家榜":
		qidiantype = crawler.NOVEL_PUBLIC_AUTHOR_RANKING
	case "月票榜VIP新作":
		qidiantype = crawler.NOVEL_VIP_RANKING
	}
	// 根据具体类型获取起点排行榜数据
	var rq = crawler.NewQiDianRankingReptile(qidiantype)
	yueps := rq.GetRanking()
	ctx.JSON(200, yueps)

}
