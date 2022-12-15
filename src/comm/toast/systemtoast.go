package toast

import (
	"log"

	"github.com/go-toast/toast"
)

// golang 通过PowerShell脚本调用系统通知
func SystemToast(title, msg string) {
	notification := toast.Notification{
		AppID:   "Microsoft.Windows.Shell.RunDialog",
		Title:   title,
		Message: msg,
		Icon:    "", // 文件必须存在
		Actions: []toast.Action{
			{"protocol", "点击查看详情", "https://www.baidu.com/"},
		},
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}
