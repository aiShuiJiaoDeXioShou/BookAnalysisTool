package crawler

// 爬取笔趣阁里面的书籍信息，爬取书籍信息，爬取所有章节

import (
	"BookAnalysisTool/src/comm/strtools"
	"BookAnalysisTool/src/parse/comm"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

type BiqupaManagement struct {
}

func NewBiqupaManagement() *BiqupaManagement {

	return &BiqupaManagement{}
}

// 这个模块是常用网址用不了了，直接去笔趣阁爬
func (biqu *BiqupaManagement) BiqupaBook(name string) (books []map[string]comm.Book) {
	var url = fmt.Sprintf("http://www.b5200.org/modules/article/search.php?searchkey='%v'", name)
	// 新建一个爬虫
	var c = colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36"),
		colly.AllowedDomains("http://www.b5200.org/"),
	)
	// 这个用于获取章节的信息
	chapterC := c.Clone()

	extensions.RandomUserAgent(c)
	c.OnHTML("#hotcontent > table > tbody > tr:not(:first-child)", func(h *colly.HTMLElement) {
		h.DOM.Each(func(i int, s *goquery.Selection) {
			var book comm.Book
			var mymap = make(map[string]comm.Book)
			title := s.Find("td:nth-child(1)").Text()
			book.Author = s.Find("td:nth-child(3)").Text()
			book.Title = title
			book.UpdateTime = s.Find("td:nth-child(5)").Text()
			book.State = s.Find("td:nth-child(6)").Text()
			book.BookUrl, _ = s.Find("td:nth-child(1)").Find("a").Attr("href")
			mymap[title] = book
			books = append(books, mymap)
			chapterC.Visit(book.BookUrl)
		})
	})
	// 获取所有的章节链接
	chapterC.OnHTML("#list > dl > *", func(h *colly.HTMLElement) {
		h.DOM.Each(func(i int, s *goquery.Selection) {
			val, _ := s.Find("a").Attr("href")
			log.Println(val)
		})
	})
	c.Visit(url)
	c.Wait()

	return
}

// 打开指定书籍的url，获取所有的章节链接，然后爬取
func (biqu *BiqupaManagement) BiqupaChapterUrls(url string) []*comm.Chapter {
	d := newGrabQuery(url)
	s := d.Find("#list > dl > *")

	// 抓取章节名称，和链接
	var b = false
	// 章节链表
	var chapters []*comm.Chapter

	// 爬取章节
	s.Each(func(i int, s *goquery.Selection) {
		if strings.Contains(s.Text(), "正文") {
			b = true
		}
		if b {
			val, _ := s.Find("a").Attr("href")
			if val != "" {
				var chapter comm.Chapter
				// 寻找章节标题
				title := s.Text()
				chapter.Title = title
				chapter.BiQuUrl = val
				chapters = append(chapters, &chapter)
			}
		}
	})

	return chapters
}

// 爬取指定章节链接的内容
func (biqu *BiqupaManagement) BiqupaChapterContents(chapters []*comm.Chapter) {
	// 根据章节列表爬取内容
	for index, chapter := range chapters {
		if index%10 == 0 {
			fmt.Println("暂停中...")
			time.Sleep(4 * time.Second)
		}
		fmt.Println("正在爬取", chapter.Title, "...")
		text_d := newGrabQuery(chapter.BiQuUrl)
		// 寻找它的一个内容
		content := text_d.Find("#content").Text()
		chapter.Contexnt = content
		index++
		fmt.Println(chapter.Title, "爬取结束...")
	}
}

// 新建一个goQuery爬虫
func newGrabQuery(url string) *goquery.Document {
	// 使用http请求，爬取页面数据
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// 将页面翻译成utf-8
	r, _ := strtools.DecodeHTMLBody(res.Body, "")
	d, _ := goquery.NewDocumentFromReader(r)
	return d
}
