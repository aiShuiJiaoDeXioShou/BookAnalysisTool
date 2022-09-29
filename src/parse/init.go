package parse

import (
	"BookAnalysisTool/src/parse/download"
	"BookAnalysisTool/src/parse/paragraphparser"
	"bytes"
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	cumulativeTime time.Duration
)

func OpenWorkSpaceParse() {

}

func SrcFlagParse(srcFlag string) {
	tpp := paragraphparser.NewParagraphParser(srcFlag)
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

// 这个只有是命令行应用才能使用
func WatchFlagParse() {
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

func SearchBookParse(searchBook string) {
	sb := download.SearchBookNameUrl(searchBook)
	// sb 转json字符串
	b, err := json.Marshal(sb)
	if err != nil {
		log.Panic(err.Error())
	}
	var str bytes.Buffer
	json.Indent(&str, b, " ", "    ")
	log.Println(str.String())
}

func DownTermParse(downTerm string, dwonUrl string) {
	log.Println("下载开始...")
	download.DownloadBook(downTerm, dwonUrl)
	log.Println("下载结束!")
}

func DownTermAllParse(downTerm string, dwonUrl string) {
	log.Println("下载开始...")
	download.DownloadAllBook(downTerm, dwonUrl)
	log.Println("下载结束!")
}

// 扫描指定文件夹,判断最近修改的文件
func ScanDir(dir string) {

}
