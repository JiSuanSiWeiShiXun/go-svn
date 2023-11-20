package svn

// SVNRepo svn repository information
type SVNRepo struct {
	URL    string
	User   string
	Passwd string
}

// SVNOption parameters encapsulated in creation of an SVNClient
type SVNOption struct {
	SVNPath string // svn path
	SVNRepo
}

// SVNClient handle for svn oprations
type SVNClient struct {
	SVNOption
}

// SVNBlame encapsulate parameters provided by `svn blame`
type SVNBlame struct {
	Order    int    `desc:"行号"`
	Revision string `desc:"版本号"`
	Author   string `desc:"提交人"`
	Date     string `desc:"最近一次修改日期"`
	Content  string `desc:"内容"`
}
