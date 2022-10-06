package commandgui

import (
	"BookAnalysisTool/src/gui"
	"BookAnalysisTool/src/parse"
	"BookAnalysisTool/src/parse/crawler"
	"flag"
)

var (
	srcFlag       *string
	watchFlag     *bool
	searchBook    *string
	downTerm      *string
	downTermAll   *string
	downBookBiqu  *string
	d             *string
	openWorkSpace *string
	guibool       *bool
	web           *bool
	timing        *int64
)

func CmdParse() {
	srcFlag = flag.String("file", "", "要解析的文件路径")
	// 开启工作区监听
	watchFlag = flag.Bool("watch", false, "是否开启工作区监听？默认为false")

	// 搜索指定名称的书籍
	searchBook = flag.String("search", "", "搜索指定名称的书籍")

	// 搜索指定名称的书籍并下载
	downTerm = flag.String("down-book", "", "搜索指定名称的书籍并下载")

	// 搜索指定名称相关的书籍并下载
	downTermAll = flag.String("down-all-book", "", "搜索指定名称相关的书籍并下载")

	// 如果传统书籍下载有问题，直接切换到笔趣阁去下载
	downBookBiqu = flag.String("down-book-biqu", "", "如果传统书籍下载有问题，直接切换到笔趣阁去下载")

	// 开启GUI动态网页
	guibool = flag.Bool("gui", false, "开启图形化界面")

	// 开启Web服务端
	web = flag.Bool("web-service", false, "不打开浏览器，开启服务器")

	// 开启定时任务并且开启web服务器
	timing = flag.Int64("timing", 0, "定时任务，默认是零的话就不开启")

	// 指定下载位置
	d = flag.String(
		"d",
		"C:/Users/Public/Documents/小说/",
		`指定文件下载位置
		--默认下载地址C:/Users/Public/Documents/小说/`)

	// 打开工作区间
	openWorkSpace = flag.String("openWorkSpace", "C:/Users/Public/Documents/", "打开工作区间，进行创作")

	flag.Parse()
	if *srcFlag != "" {
		parse.SrcFlagParse(*srcFlag)
	}

	if *watchFlag {
		parse.WatchFlagParse()
	}

	if *searchBook != "" {
		parse.SearchBookParse(*searchBook)
	}

	if *downTerm != "" {
		parse.DownTermParse(*downTerm, *d)
	}

	// 下载所有相关的
	if *downTermAll != "" {
		parse.DownTermAllParse(*downTermAll, *d)
	}

	// 笔趣阁下载指定相关
	if *downBookBiqu != "" {
		bm := crawler.NewBiqupaManagement()
		bm.BiqupaBookToTxt(*downBookBiqu, *d+*downBookBiqu)
	}

	// 打开工作区间
	if *openWorkSpace != "" {
		parse.OpenWorkSpaceParse()
	}

	// 打开gui
	if *guibool {
		gui.Exec()
	}

	// 只是单纯的启用web Api接口
	if *web {
		gui.ApiStart()
	}

	// 开启定时任务
	if *timing > 0 {

	}
}
