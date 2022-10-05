package download

import (
	"BookAnalysisTool/src/parse/comm"
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
	// 放置下载资源的地方
	downloadPath = "C:\\Users\\Public\\Documents\\小说\\"
)

// 搜索指定图书的下载链接
func SearchBookNameUrl(name string) (searchBooks []*comm.SearchBook) {
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
			var searchBook comm.SearchBook
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
	reader := bufio.NewReader(res.Body)

	// 判断downloadPath文件夹是否存在,不存在则创建
	_, filePathIsExistErr := os.Lstat(downloadPath)
	if filePathIsExistErr != nil {
		fmt.Println("A error occurred!")
		// 创建文件夹
		err2 := os.MkdirAll(downloadPath, 0777)
		if os.IsNotExist(err2) {
			fmt.Println("创建文件夹失败!")
		}
	}

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
func DownloadAllBook(name string, downPath string) {
	var wg sync.WaitGroup
	sb := SearchBookNameUrl(name)
	wg.Add(len(sb))
	for i := 0; i < len(sb); i++ {
		go func(i int) {
			OpenDownload(sb[i].DownloadUrl, downPath)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Downloaded Done!")
}

// 下载指定书籍
func DownloadBook(name string, downPath string) {
	sb := SearchBookNameUrl(name)
	for _, s := range sb {
		if s.Title == name {
			OpenDownload(s.DownloadUrl, downPath)
			return
		}
	}
	OpenDownload(sb[0].DownloadUrl, downPath)
}

// 起点数据榜
// 菠萝包
// 刺猬猫
