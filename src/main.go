package main

import (
	"flag"
	"io/ioutil"
	"log"
)

var (
	srcFlag *string
)

func srcFlagParse() {
	// 读取src文件的内容
	if *srcFlag != "" {
		b, err := ioutil.ReadFile(*srcFlag)
		if err != nil {
			log.Fatal(err)
		}
		// 将读取的内容转换为string
		s := string(b)
		// 输出文件内容
		log.Println(s)
	}
}

func main() {
	srcFlag = flag.String("file", "", "要解析的文件路径")
	flag.Parse()
	srcFlagParse()
}
