package main

import (
	"BookAnalysisTool/src/comm/strtools"
	"BookAnalysisTool/src/service/comm"
	"BookAnalysisTool/src/service/crawler"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func TestBaiDu(t *testing.T) {
	b := crawler.NewBaiDuBookHotManagement()
	books := b.BaiDuCrawler("全部类型")
	t.Log(books)
}

func TestQiDian(t *testing.T) {
	bm := crawler.NewBiqupaManagement()
	b := bm.BiqupaBook("道诡异仙")
	t.Log(b)
}

func TestBiQU(t *testing.T) {
	res, err := http.Get("http://www.b5200.org/0_844/636719.html")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	r, _ := strtools.DecodeHTMLBody(res.Body, "")
	d, _ := goquery.NewDocumentFromReader(r)
	s := d.Find("#content").Text()
	t.Log(s)

}

func TestBiquPa(t *testing.T) {
	BiquPa()
}

// 抓取章节的内容
func grabText() {

}

func BiquPa() {
	d := newGrabQuery("http://www.b5200.org/52_52542/")
	s := d.Find("#list > dl > *")

	// 抓取章节名称，和链接
	var b = false
	// 章节链表
	var chapters []comm.Chapter

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
				chapters = append(chapters, chapter)
			}
		}
	})

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

	json, _ := strtools.MyJsonFormat(chapters)
	ioutil.WriteFile("data.json", []byte(json), 0766)

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

func TestBiqu(t *testing.T) {
	bm := crawler.NewBiqupaManagement()
	books := bm.BiqupaBook("一世之尊")
	json, _ := strtools.MyJsonFormat(books)
	ioutil.WriteFile("一世之尊.json", []byte(json), 0766)
}

func Test999(t *testing.T) {
	escapeUrl := url.QueryEscape("诡秘之主")
	fmt.Println("编码:", escapeUrl)

}
