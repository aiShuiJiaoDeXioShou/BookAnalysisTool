package strtools

import "strings"

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
