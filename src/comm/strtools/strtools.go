package strtools

import (
	"bytes"
	"encoding/json"
	"strings"
)

// 将整段文字转化为一个自然段数组,并且去掉单行空格
func SplitTextParagraphs(text string) []string {
	s := strings.Split(string(text), "\n")
	for index, str := range s {
		if len(strings.TrimSpace(str)) <= 0 {
			s = append(s[:index], s[index+1:]...)
		}
	}

	return s
}

// 去掉文件的后缀名
func RmSuffix(str string) (strname string) {
	subI := strings.LastIndex(str, "/")
	if subI != -1 {
		strname = str[subI+1:]
	}

	subI2 := strings.LastIndex(str, "\\")
	if subI2 != -1 {
		strname = str[subI2:]
	}
	i := strings.LastIndex(strname, ".")
	strname = strname[:i]

	return
}

// 将对象转化为json对象,并且格式化该对象
func MyJsonFormat(project any) (string, error) {
	// sb 转json字符串
	b, err := json.Marshal(project)
	var str bytes.Buffer
	json.Indent(&str, b, " ", "    ")
	return str.String(), err
}
