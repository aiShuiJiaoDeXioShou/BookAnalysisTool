package parse

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"text_processing/src/comm/strtools"
	"text_processing/src/config"
)

type TextParseProject struct {
	Path    string   `json:"path"`
	Content []string `json:"content"`
	Text    string   `json:"text"`
	Size    []byte   `json:"size"`
}

// 初始化函数
func NewTextParseProject(path string) *TextParseProject {
	var text string
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic(err.Error())
	}
	text = string(b)
	content := strings.Fields(text)

	return &TextParseProject{
		Path:    path,
		Content: content,
		Size:    b,
		Text:    text,
	}
}

// 统计一共多少个字(只算汉字)
func (t *TextParseProject) Count() float64 {

	return float64(len(t.Text) / 3)
}

// 统计某个字的出现次数
func (t *TextParseProject) CountWord(word string) float64 {

	return float64(strings.Count(t.Text, word))
}

// 统计某个字的出现频率
func (t *TextParseProject) CountWordFrequency(word string) float64 {
	i := t.CountWord(word) * float64(len(word)/3)
	total := t.Count()
	return i / total
}

// 根据文件名获取书名
func (t *TextParseProject) PraseBookName() string {
	// 去掉后缀名
	s := strtools.RmSuffix(t.Path)
	// 对s进行预处理,检查是否有违规字段
	s2, err := t.SearchViolateStr(s)
	if err != nil {
		log.Printf("检测到有违规字的存在-->\n%s", err.Error())
	}
	log.Printf("你将得到该书名:%s", s2)
	return s2
}

// 根据书名查找作者信息，查找书籍简介
func (t *TextParseProject) FindAuthor() string {

	return ""
}

// 检查是否有违规字段,返回不合法的名称,并且修改当前不合法字段
func (t *TextParseProject) SearchViolateStr(search string) (string, error) {
	s := config.ViolateStr
	for _, str := range s {
		i := strings.Index(search, str)
		if i != -1 {
			prestr := search[:i] + "***" + search[i+1:]
			newStr := strings.Trim(prestr, str)
			return newStr, errors.New(fmt.Sprintf("该[%s]字符串不能使用", str))
		}
	}

	return search, nil
}

// 查找某个单词或者词语的意思
func (t *TextParseProject) FindWordMeaning(word string) string {

	return ""
}
