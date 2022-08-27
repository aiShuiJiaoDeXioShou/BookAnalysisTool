package main

import (
	"BookAnalysisTool/src/parse/download"
	"BookAnalysisTool/src/parse/paragraphparser"
	"bytes"
	"encoding/json"
	"flag"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	srcFlag        *string
	watchFlag      *bool
	zushe          chan bool
	cumulativeTime time.Duration
	searchBook     *string
	downTerm       *string
	downTermAll    *string
	d              *string
)

func main() {
	srcFlag = flag.String("file", "", "要解析的文件路径")
	// 开启工作区监听
	watchFlag = flag.Bool("watch", false, "是否开启工作区监听？默认为false")

	// 搜索指定名称的书籍
	searchBook = flag.String("search", "", "搜索指定名称的书籍")

	// 搜索指定名称的书籍并下载
	downTerm = flag.String("down-term", "", "搜索指定名称的书籍并下载")

	// 搜索指定名称相关的书籍并下载
	downTermAll = flag.String("down-all-term", "", "搜索指定名称相关的书籍并下载")

	// 指定下载位置
	d = flag.String("d", "C:/Users/Public/Documents/小说", "指定文件下载位置")
	commParse()
}

func commParse() {
	flag.Parse()
	if *srcFlag != "" {
		srcFlagParse()
	}

	if *watchFlag {
		watchFlagParse()
	}

	if *searchBook != "" {
		searchBookParse()
	}

	if *downTerm != "" {
		downTermParse()
	}

	// 下载所有相关的
	if *downTermAll != "" {
		downTermAllParse()
	}
}

func srcFlagParse() {
	tpp := paragraphparser.NewParagraphParser(*srcFlag)
	b, err := json.Marshal(tpp)
	if err != nil {
		log.Panic(err.Error())
	}
	// 打印信息
	log.Println(string(b))
	// golang 写入文件
	err2 := ioutil.WriteFile("build/data.json", b, fs.ModeAppend)
	if err2 != nil {
		log.Panic(err2.Error())
	}
}

func watchFlagParse() {
	go func() {
		for {
			time.Sleep(time.Second)
			cumulativeTime += time.Second
		}
	}()
	log.Println("工作区监听已开启:")
	log.Println("工作区默认将在( D:\\workspace\\write ) 下开启...")
	log.Println("按任意键结束监听...")
	b := make([]byte, 1)
	os.Stdin.Read(b)
	// 程序退出的时候执行一些任务
	defer func() {
		log.Printf("小禾(表示):太短了【%s】", cumulativeTime.String())
	}()

}

func searchBookParse() {
	sb := download.SearchBookNameUrl(*searchBook)
	// sb 转json字符串
	b, err := json.Marshal(sb)
	if err != nil {
		log.Panic(err.Error())
	}
	var str bytes.Buffer
	json.Indent(&str, b, " ", "    ")
	log.Println(str.String())
}

func downTermParse() {
	log.Println("下载开始...")
	download.DownloadBook(*downTerm, *d)
	log.Println("下载结束!")
}

func downTermAllParse() {
	log.Println("下载开始...")
	download.DownloadAllBook(*downTerm, *d)
	log.Println("下载结束!")
}

// 扫描指定文件夹,判断最近修改的文件
func scanDir(dir string) {

}
