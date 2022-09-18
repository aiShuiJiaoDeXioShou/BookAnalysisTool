package gui

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"github.com/gin-gonic/gin"
)

// 首先它会打开浏览器
func Exec() {
	start()
}

// 打开浏览器并跳转到首页
func toIndex() {
	// 如果服务器启动成功则打开首页
	// 无GUI调用
	if strings.Contains("windows", runtime.GOOS) {
		cmd := exec.Command(`cmd`, `/c`, `start`, `http://localhost:8031/`)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		err := cmd.Start()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("如果打开首页失败，看是不是浏览器版本太低了，请升级浏览器！")
		fmt.Println("将这段链接直接粘贴到浏览器导航栏当中，即可跳转到首页！")
		fmt.Println("http://localhost:8031/")
	} else {
		cmd := exec.Command(`xdg-open`, `http://localhost:8031/`).Start()
		// 如果启动linux出错那么肯定是macos
		if cmd != nil {
			exec.Command(`open`, `http://localhost:8031/`).Start()
		}
	}
}

func start() {
	e :=gin.Default()
	// 资源映射配置
	static(e)
	// 这个是Api配置
	api(e)
	toIndex()
	err := e.Run(":8031")
	if err != nil {
		panic("滴滴，端口被占领，请联系程序员更换端口号！")
	}
}

// api配置
func api(e *gin.Engine) {
	e.GET("/",func(ctx *gin.Context) {
		ctx.HTML(200,"index.html",nil)
	})
}

// 解决静态资源问题,办法就是给常用的全部映射
func static(e *gin.Engine) {
	//这个是模板目录
	e.LoadHTMLGlob("views/**/*")
	e.Static("/static","static")
	e.Static("/dist", "./dist")
	e.Static("/css", "./dist/css")
	e.Static("/js", "./dist/js")
	e.StaticFile("/favicon.ico", "./dist/favicon.ico")
}