package crawler

type GitHubCrawler struct {
}

func NewGitHubCrawler() *GitHubCrawler {
	return &GitHubCrawler{}
}

// github 仓库的信息详情
type Repo struct {
	Name    string   // 仓库名
	Author  string   // 作者名
	Link    string   // 链接
	Desc    string   // 描述
	Lang    string   // 语言
	Stars   int      // 星数
	Forks   int      // fork 数
	Add     int      // 周期内新增
	BuiltBy []string // 贡献值 avatar img 链接
}
