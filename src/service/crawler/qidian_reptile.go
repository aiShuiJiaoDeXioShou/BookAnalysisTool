/**
* @作者: 林河
* @date: 2022/6/27
* @describe:
**/
package crawler

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// 起点中文网的爬虫,定义各个榜单的常量
const (
	// 月票榜
	NOVEL_RANKING = "body > div.wrap > div.rank-box.box-center.cf > div.main-content-wrap.fl > div.rank-body > div > div:nth-child(1) > div > ul"

	// 畅销榜
	NOVEL_SALES_RANKING = "body > div.wrap > div.rank-box.box-center.cf > div.main-content-wrap.fl > div.rank-body > div > div:nth-child(2) > div > ul"

	// 阅读指数
	NOVEL_READ_RANKING = "body > div.wrap > div.rank-box.box-center.cf > div.main-content-wrap.fl > div.rank-body > div > div:nth-child(3) > div.tab-list > div > ul"

	// 推荐榜
	NOVEL_RECOMMEND_RANKING = "body > div.wrap > div.rank-box.box-center.cf > div.main-content-wrap.fl > div.rank-body > div > div:nth-child(4) > div.tab-list > div:nth-child(2) > ul"

	// 收藏榜
	NOVEL_COLLECT_RANKING = "body > div.wrap > div.rank-box.box-center.cf > div.main-content-wrap.fl > div.rank-body > div > div:nth-child(5) > div > ul"

	// 签约作者榜
	NOVEL_NEW_AUTHOR_RANKING = "body > div.wrap > div.rank-box.box-center.cf > div.main-content-wrap.fl > div.rank-body > div > div:nth-child(6) > div > ul"

	// 公共作家榜
	NOVEL_PUBLIC_AUTHOR_RANKING = "body > div.wrap > div.rank-box.box-center.cf > div.main-content-wrap.fl > div.rank-body > div > div:nth-child(7) > div > ul"

	// 月票榜VIP新作
	NOVEL_VIP_RANKING = "body > div.wrap > div.rank-box.box-center.cf > div.main-content-wrap.fl > div.rank-body > div > div:nth-child(8) > div > ul"
)

// 爬取月票榜
// 起点中文网月票结构体
type QidianRanking struct {
	// 排名
	SequenceID int `json:"sequenceId"`
	// 小说名称
	NovelName string `json:"novelName"`
	// 月票数量
	MonthTicket int `json:"monthTicket"`
	// 小说链接
	NovelLink string `json:"novelLink"`
}

// 爬虫管理器哦
type QiDianRankingReptile struct {
	// 要抓取的排行根源目录
	root string
}

func NewQiDianRankingReptile(root string) *QiDianRankingReptile {

	return &QiDianRankingReptile{
		root: root,
	}
}

// 写一个爬取起点月票板的爬虫
func (q *QiDianRankingReptile) GetRanking() *[]QidianRanking {

	// 月票榜数据
	var qidians []QidianRanking

	// 这里是爬虫的参数
	userAgent := colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Safari/537.36 Edg/103.0.1264.37")
	c := colly.NewCollector(userAgent)
	// 设置爬虫的请求头

	// 解析月票榜
	c.OnHTML(q.root, func(e *colly.HTMLElement) {

		e.ForEach("li[data-rid]", func(i int, item *colly.HTMLElement) {
			var qidian QidianRanking

			item.ForEach("div", func(i int, item *colly.HTMLElement) {
				q.解析排行榜的内层html解析器(i, item, &qidian)
			})
			qidians = append(qidians, qidian)
		})

	})

	err := c.Visit("https://www.qidian.com/rank/")
	if err != nil {
		log.Println(err.Error())
	}

	return &qidians
}

func (q *QiDianRankingReptile) 第一个书籍html解析器(i int, item *colly.HTMLElement, qidian *QidianRanking) {
	qidian.SequenceID = 1
	selectOrId := ".book-wrap.cf>.book-info.fl"
	qidian.NovelName = item.ChildText(selectOrId + ">h2>a")
	monthTicket := item.ChildText(selectOrId + " p.digital > em")

	num := strings.Replace(monthTicket, "月票数量:", "", -1)

	monthTicketInt, err := strconv.Atoi(strings.TrimSpace(num))

	if err != nil {
		log.Println(err.Error())
	}

	qidian.MonthTicket = monthTicketInt

}

func (q *QiDianRankingReptile) 解析排行榜的内层html解析器(i int, item *colly.HTMLElement, qidian *QidianRanking) {
	selectOrId := ".book-wrap.cf>.book-info.fl"
	text := item.ChildText(selectOrId)
	if text != "" {
		q.第一个书籍html解析器(i, item, qidian)
	} else {
		if i == 0 {
			// 小说排行
			novelName := item.Text
			// 将novelName强转为int类型
			novelNameInt, _ := strconv.Atoi(novelName)
			qidian.SequenceID = novelNameInt

		} else if i == 1 {

			href := item.ChildAttr("a", "href")
			qidian.NovelLink = href
			name := item.ChildText("a")
			qidian.NovelName = name
			total := item.ChildText("i[class=total]")
			totalInt, _ := strconv.Atoi(total)
			if qidian.MonthTicket != 0 {
				return
			}
			qidian.MonthTicket = totalInt
		}
	}

}
