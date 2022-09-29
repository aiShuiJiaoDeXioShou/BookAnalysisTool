package crawler

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

type BaiDuBookHotSearch struct {
	Image     string `json:"image"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Type      string `json:"type"`
	Details   string `json:"details"`
	HotNumber string `json:"hotnumber"`
	Number    string `json:"number"`
}

// 百度热搜榜，可见数据
func BaiDuCrawler(typestr string) *[]BaiDuBookHotSearch {
	if typestr == "" {
		typestr = "全部类型"
	}
	var queryId = "#sanRoot > main > div.container.right-container_2EFJr > div > div:nth-child(2) > .category-wrap_iQLoo"
	var hotBooks []BaiDuBookHotSearch
	c := colly.NewCollector(
		// colly.AllowedDomains("https://top.baidu.com/board?platform=pc&tab=novel&tag={'category':'全部类型'}"),
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36"),
	)
	// 添加相关网站对应的Cookie值
	c.OnRequest(func(r *colly.Request) {

	})
	// 将自己伪装成普通浏览器
	extensions.RandomUserAgent(c)
	c.OnHTML(queryId, func(h *colly.HTMLElement) {
		h.DOM.Each(func(i int, s *goquery.Selection) {
			var book BaiDuBookHotSearch
			// 取到封面数据
			imgUrl, _ := s.Find(".img-wrapper_29V76>img").Attr("src")
			book.Image = imgUrl
			// 取到序号
			num := s.Find(".img-wrapper_29V76>index_1Ew5p").Text()
			book.Number = num
			// 取到标题
			title := s.Find(".content_1YWBm>a").Text()
			book.Title = title
			// 作者
			author := s.Find(".content_1YWBm>div:nth-child(1)").Text()
			book.Author = author
			// 类型
			booktype := s.Find(".content_1YWBm>div:nth-child(2)").Text()
			book.Type = booktype
			details := s.Find(".content_1YWBm>div:nth-child(3)").Text()
			book.Details = details
			// 拿到热搜指数
			hotNumber := s.Find(".trend_2RttY>div:nth-child(2)").Text()
			book.HotNumber = hotNumber

			hotBooks = append(hotBooks, book)
		})
	})
	// 来开始爬bilibili
	c.Visit("https://top.baidu.com/board?platform=pc&tab=novel&tag={'category':'" + typestr + "'}")
	// 这个函数要放在最后面
	c.Wait()
	return &hotBooks
}
