package crawler

// 爬取笔趣阁里面的书籍信息，爬取书籍信息，爬取所有章节

import (
	"BookAnalysisTool/src/comm/strtools"
	"BookAnalysisTool/src/parse/comm"
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type BiqupaManagement struct {
}

func NewBiqupaManagement() *BiqupaManagement {

	return &BiqupaManagement{}
}

// 在笔趣阁里面获取指定书籍所有的内容
func (biqu *BiqupaManagement) BiqupaBook(bookname string) comm.Book {
	// 获取相关的所有书籍
	m := biqu.BiqupaBookUrls(bookname)
	book := m[bookname]
	// 根据书籍链接获取所有章节链接
	chaUrls := biqu.BiqupaChapterUrls(book.BookUrl)
	// 将这些章节填充内容
	biqu.BiqupaChapterContents(chaUrls)
	book.Chapters = chaUrls
	return book
}

// 将指定书籍的内容转化为指定的txt文档
func (biqu *BiqupaManagement) BiqupaBookToTxt(name, url string) {
	b := biqu.BiqupaBook(name)
	f, err := os.OpenFile(url+".txt", os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer f.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(f)
	write.WriteString("\n" + b.Title + "\n")
	write.WriteString(b.Author + "\n")
	for _, c := range b.Chapters {
		write.WriteString("\n" + c.Title + "\n")
		write.WriteString("\n" + c.Contexnt + "\n")
	}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}

// 这个模块是常用网址用不了了，直接去笔趣阁爬，获取所有书籍链接
func (biqu *BiqupaManagement) BiqupaBookUrls(name string) map[string]comm.Book {
	var url = fmt.Sprintf("http://www.b5200.org/modules/article/search.php?searchkey=%v", url.QueryEscape(name))
	// 存放数据的链表
	var mymap = make(map[string]comm.Book)
	// 使用http请求，爬取页面数据
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	d, _ := goquery.NewDocumentFromReader(res.Body)
	d.Find("#hotcontent > table > tbody > tr:not(:first-child)").Each(func(i int, s *goquery.Selection) {
		var book comm.Book
		title := strtools.ConvertToString(s.Find("td:nth-child(1)").Text(), "gbk", "utf-8")
		book.Author = strtools.ConvertToString(s.Find("td:nth-child(3)").Text(), "gbk", "utf-8")
		book.Title = title
		book.UpdateTime = strtools.ConvertToString(s.Find("td:nth-child(5)").Text(), "gbk", "utf-8")
		book.State = strtools.ConvertToString(s.Find("td:nth-child(6)").Text(), "gbk", "utf-8")
		a_txt, _ := s.Find("td:nth-child(1)").Find("a").Attr("href")
		book.BookUrl = strtools.ConvertToString(a_txt, "gbk", "utf-8")
		// 获取书籍封面数据
		// book.BookWebUrl = strtools.ConvertToString(getBookImage(book.BookUrl), "gbk", "utf-8")
		mymap[title] = book
	})
	d.Clone()
	return mymap
}

func getBookImage(url string) string {
	// 使用http请求，爬取页面数据
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	d, _ := goquery.NewDocumentFromReader(res.Body)
	val, exists := d.Find("#fmimg").Find("img").Attr("src")
	if !exists {
		log.Println("封面不存在!")
	}
	d.Clone()
	return val
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
		var content string
		var duanNum = 0
		text_d.Find("#content>p").Each(func(i int, s *goquery.Selection) {
			duan := s.Text()
			content += "\n\n" + duan
			duanNum++
		})
		// 如果小于2说明爬取这本书的时候有bug，我们把所有空格换成\n\n\t\t
		if duanNum < 2 {
			content, _ = strtools.ReplaceStringByRegex(content, "“", "“\n\t\t")
			content, _ = strtools.ReplaceStringByRegex(content, "。", "。\n\t\t")
		}
		chapter.Contexnt = content
		index++
		fmt.Println(chapter.Title, "爬取结束...")
	}
	fmt.Println("书籍爬取完成！")
}

// 爬取单个章节内容
func (biqu *BiqupaManagement) BiqupaOneChapterContent(chapter *comm.Chapter) {
	fmt.Println("正在爬取", chapter.Title, "...")
	text_d := newGrabQuery(chapter.BiQuUrl)
	// 寻找它的一个内容
	var content string
	var duanNum = 0
	text_d.Find("#content>p").Each(func(i int, s *goquery.Selection) {
		duan := s.Text()
		content += "\n\n" + duan
		duanNum++
	})
	// 如果小于2说明爬取这本书的时候有bug，我们把所有空格换成\n\n\t\t
	if duanNum < 2 {
		content, _ = strtools.ReplaceStringByRegex(content, "“", "“\n\t\t")
		content, _ = strtools.ReplaceStringByRegex(content, "。", "。\n\t\t")
	}
	chapter.Contexnt = content
	fmt.Println(chapter.Title, "爬取结束...")
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
