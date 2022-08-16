package main

import (
	"encoding/json"
	"flag"
	"io/fs"
	"io/ioutil"
	"log"
	"text_processing/src/parse"
)

var (
	srcFlag *string
)

func srcFlagParse() {
	// 读取src文件的内容
	if *srcFlag != "" {
		tpp := parse.NewTextParseProject(*srcFlag)
		b, err := json.Marshal(tpp)
		if err != nil {
			log.Panic(err.Error())
		}

		// golang 写入文件
		err2 := ioutil.WriteFile("build/data.json", b, fs.ModeAppend)
		if err2 != nil {
			log.Panic(err2.Error())
		}

	}
}

func main() {
	srcFlag = flag.String("file", "", "要解析的文件路径")
	flag.Parse()
	srcFlagParse()
}
