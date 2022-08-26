package parse

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

const (
	// 歌书网的搜索链接
	baseUrl = "http://www.gebiqu.com/modules/article/search.php?searchkey="
)

type SearchBook struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Author      string `json:"author"`
	Number      string `json:"number"`
	State       string `json:"state"`
	UpdateTime  int64  `json:"update_time"`
	DownloadUrl string `json:"download_url"`
}

// 搜索指定图书的下载链接
func SearchBookNameUrl(name string) (searchBooks []*SearchBook) {
	res, err := http.Get(baseUrl + name)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	nrs := doc.Find("#nr")
	var wgG sync.WaitGroup
	wgG.Add(nrs.Length())
	nrs.Each(func(i int, s *goquery.Selection) {
		go func() {
			var searchBook SearchBook
			var wg sync.WaitGroup

			wg.Add(2)
			var tds = s.Children()
			// 获取标题和URL信息
			titleOrUrl := tds.Eq(0)
			searchBook.Title = titleOrUrl.Text()
			val, exists := titleOrUrl.Children().Attr("href")
			if !exists {
				return
			}
			searchBook.URL = val

			go func() {
				// 获取作者信息
				searchBook.Author = tds.Eq(2).Text()
				// 获取字数
				searchBook.Number = tds.Eq(3).Text()
				// 获取时间
				// 将字符串转化为int64
				number, numberToStringErr := strconv.Atoi(tds.Eq(4).Text())
				if numberToStringErr != nil {
					log.Println(numberToStringErr)
				}
				searchBook.UpdateTime = int64(number)
				// 获取状态
				searchBook.State = tds.Eq(5).Text()
				wg.Done()
			}()

			go func() {
				// 获取下载链接
				// searchBook.DownloadUrl
				d, downUrlErr := goquery.NewDocument(val)
				if downUrlErr != nil {
					log.Println(downUrlErr.Error())
				}
				downUrl := d.Find("#hotcontent > div.l > div > table > tbody > tr > td:nth-child(2) > div:nth-child(5) > a:nth-child(2)")
				searchBook.DownloadUrl, exists = downUrl.Attr("href")
				if !exists {
					log.Println("获取下载链接失败", searchBook.Title)
				}
				wg.Done()
			}()

			wg.Wait()
			// 将searchBook添加到searchBooks
			searchBooks = append(searchBooks, &searchBook)
			wgG.Done()
		}()

	})
	wgG.Wait()
	return
}

// 将获取的下载链接下载指定书籍到指定文件夹
func OpenDownload(filename string, downloadPath string) {
	fileName := path.Base(filename)

	res, err := http.Get(filename)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	file, err := os.Create(downloadPath + fileName)
	if err != nil {
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d", written)
}

// 下载歌书网所有有关斗破苍穹的图书
func DownloadBook(name string) {
	var wg sync.WaitGroup
	sb := SearchBookNameUrl(name)
	wg.Add(len(sb))
	for i := 0; i < len(sb); i++ {
		go func(i int) {
			OpenDownload(sb[i].DownloadUrl, "C:\\Users\\28322\\Desktop\\text_processing\\src\\res\\")
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Downloaded Done!")
}
