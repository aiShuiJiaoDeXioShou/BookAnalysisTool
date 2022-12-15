package main

import (
	"BookAnalysisTool/src/service/crawler"
	"log"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/fsnotify/fsnotify"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

// 爬虫练习biquge
func TestPanchon(t *testing.T) {
	var queryId = "#sanRoot > main > div.container.right-container_2EFJr > div > div:nth-child(2) > .category-wrap_iQLoo"
	var hotBooks []crawler.BaiDuBookHotSearch
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
			var book crawler.BaiDuBookHotSearch
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
		log.Printf("感觉没有爱情%v", hotBooks)
	})
	// 来开始爬bilibili
	c.Visit("https://top.baidu.com/board?platform=pc&tab=novel&tag={'category':'全部类型'}")
	// 这个函数要放在最后面
	c.Wait()
}

func TestTUI(t *testing.T) {
	// 使用Tui程序库

}

// 使用文件监听库
func TestFsWacth(t *testing.T) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		panic("监听失败!")
	}
	defer w.Close()

	// 启动监听事件,我们创建一个协程
	go func() {
		for {
			select {
			case event, ok := <-w.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Has(fsnotify.Write) {
					log.Println("写入的文件内容为:", event.Name)
				}
			case err, ok := <-w.Errors:
				if !ok {
					return
				}
				log.Println(err)
			}
		}
	}()
	// 添加监听的文件夹
	err = w.Add("D:\\wordtree")
	if err != nil {
		log.Fatalln(err.Error())
	}
	<-make(chan struct{})
}
