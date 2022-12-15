package comm

type Chapter struct {
	Title      string `json:"title"`
	UpdateTime string `json:"updateTime"`
	Author     string `json:"author"`
	Contexnt   string `json:"contexnt"`
	// 笔趣阁的链接
	BiQuUrl string `json:"biquurl"`
}
