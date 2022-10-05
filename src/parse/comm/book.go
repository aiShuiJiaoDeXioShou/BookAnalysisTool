package comm

type Book struct {
	Title       string    `json:"title"`
	UpdateTime  string    `json:"updateTime"`
	Author      string    `json:"author"`
	Chapters    []Chapter `json:"chapters"`
	BiQuUrl     string    `json:"biquurl"`    // 笔趣阁的链接
	BookWebUrl  string    `json:"bookweburl"` // 封面
	State       string    `json:"state"`      // 书籍的状态，连载和未连载
	BookUrl     string    `json:"bookurl"`
	ChaptersUrl []string  `json:"chaptersurl"`
}
