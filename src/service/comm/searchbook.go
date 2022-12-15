package comm

type SearchBook struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Author      string `json:"author"`
	Number      string `json:"number"`
	State       string `json:"state"`
	UpdateTime  int64  `json:"update_time"`
	DownloadUrl string `json:"download_url"`
}
