package parse

type TextParseProject struct {
	Name    string   `json:"name"`
	Path    string   `json:"path"`
	Content []string `json:"content"`
	Text    string   `json:"text"`
	Size    []byte   `json:"size"`
}

// 初始化函数
func NewTextParseProject(name string, path string, content []string, text string, size []byte) *TextParseProject {

	return &TextParseProject{
		Name:    name,
		Path:    path,
		Content: content,
		Size:    size,
		Text:    text,
	}
}

// 统计一共多少个字
func (t *TextParseProject) Count() int {

	return len(t.Text)
}

// 统计某个字的出现次数
func (t *TextParseProject) CountWord(word string) int {
	// var count int

	return 0
}

// 统计某个字的出现频率
func (t *TextParseProject) CountWordFrequency(word string) float64 {

	return 0.0
}

// 根据书名查找作者信息，查找书籍简介
func (t *TextParseProject) FindAuthor() string {

	return ""
}

// 查找某个单词或者词语的意思
func (t *TextParseProject) FindWordMeaning(word string) string {

	return ""
}
