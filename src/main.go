package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"text_processing/src/parse"
	"time"
)

var (
	srcFlag        *string
	watchFlag      *bool
	zushe          chan bool
	cumulativeTime time.Duration
)

func srcFlagParse() {
	tpp := parse.NewTextParseProject(*srcFlag)
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

	for {
		time.Sleep(time.Second)
		cumulativeTime += time.Second
	}

}

// 扫描指定文件夹,判断最近修改的文件
func scanDir(dir string) {

}

func main() {
	srcFlag = flag.String("file", "", "要解析的文件路径")
	// 开启工作区监听
	watchFlag = flag.Bool("watch", true, "是否开启工作区监听？默认为true")
	flag.Parse()
	if *srcFlag != "" {
		srcFlagParse()
	}

	if *watchFlag {
		go watchFlagParse()

		log.Println("工作区监听已开启:")
		log.Println("工作区默认将在( D:\\workspace\\write ) 下开启...")
		log.Println("1. 输入1结束监听...")
		log.Println("2. 输入2切换监听目录...")

		var a int
		_, err := fmt.Scan(&a)
		if err != nil {
			log.Panic(err.Error())
			return
		}

		if a != 9 {
			// 程序退出的时候执行一些任务,你按ctrl+C不久能到达了吗
			defer func() {
				log.Printf("小禾(表示):太短了【%s】", cumulativeTime.String())
			}()
			return
		}
	}

}
