package main

import (
	"encoding/json"
	"testing"
	"text_processing/src/parse"
)

func Test2019(t *testing.T) {
	sb := parse.SearchBookNameUrl("斗破苍穹")
	// sb 转json字符串
	b, err := json.Marshal(sb)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(b))
}

func Test2020(t *testing.T) {
	parse.OpenDownload("http://down.gebiqu.com/txt/168070/斗破苍穹.txt", "C:\\Users\\28322\\Desktop\\text_processing\\src\\res\\")
}

func Test2021(t *testing.T) {
	parse.DownloadBook("斗破苍穹")
}

func Test2022(t *testing.T) {
	parse.DownloadBook("诡秘之主")
}
