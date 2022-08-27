package main

import (
	download "BookAnalysisTool/src/parse/download"
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func Test2019(t *testing.T) {
	sb := download.SearchBookNameUrl("斗破苍穹")
	// sb 转json字符串
	b, err := json.Marshal(sb)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(b))
}

func Test2020(t *testing.T) {
	download.OpenDownload("http://down.gebiqu.com/txt/168070/斗破苍穹.txt", "C:\\Users\\28322\\Desktop\\BookAnalysisTool\\src\\res\\")
}

func Test2021(t *testing.T) {
	download.DownloadAllBook("斗破苍穹", "C:\\Users\\Public\\Documents\\小说\\")
}

func Test2022(t *testing.T) {
	download.DownloadAllBook("诡秘之主", "C:\\Users\\Public\\Documents\\小说\\")
}

func Test2023(t *testing.T) {
	sb := download.SearchBookNameUrl("诡秘之主")
	// sb 转json字符串
	b, err := json.Marshal(sb)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(b))
	download.DownloadBook("诡秘之主", "C:\\Users\\Public\\Documents\\小说\\")
}

func Test2024(t *testing.T) {
	const downloadPath = "C:\\Users\\Public\\Documents\\小说\\"
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
}

func Test2025(t *testing.T) {
	down := "http://down.gebiqu.com/txt/81731/诡秘之主.txt"
	download.OpenDownload(down, "C:/Users/Public/Documents/小说/")
}
