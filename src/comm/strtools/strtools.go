package strtools

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"strings"

	"github.com/axgle/mahonia"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/htmlindex"
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

// 转化指定编码
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func detectContentCharset(body io.Reader) string {
	r := bufio.NewReader(body)
	if data, err := r.Peek(1024); err == nil {
		if _, name, _ := charset.DetermineEncoding(data, ""); len(name) != 0 {
			return name
		}
	}

	return "utf-8"
}

// 爬虫的时候转化网页编码
func DecodeHTMLBody(body io.Reader, charset string) (io.Reader, error) {
	if charset == "" {
		charset = detectContentCharset(body)
	}

	e, err := htmlindex.Get(charset)
	if err != nil {
		return nil, err
	}

	if name, _ := htmlindex.Name(e); name != "utf-8" {
		body = e.NewDecoder().Reader(body)
	}

	return body, nil
}
