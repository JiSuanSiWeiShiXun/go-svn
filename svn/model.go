package svn

type (
	// SVNFactory factory for svn client
	SVNFactory struct{}

	SVNOption func(*SVNConfig)

	// SVNRepo svn repository information
	SVNRepo struct {
		URL    string
		User   string
		Passwd string
	}

	// SVNConfig parameters encapsulated in creation of an SVNClient
	SVNConfig struct {
		SVNPath string // svn executable path
		*SVNRepo
	}

	// SVNClient handle for svn oprations
	SVNClient struct {
		*SVNConfig
	}

	// SVNBlame encapsulate parameters provided by `svn blame`
	SVNBlame struct {
		Order    int    `desc:"行号"`
		Revision string `desc:"版本号"`
		Author   string `desc:"提交人"`
		Date     string `desc:"最近一次修改日期"`
		Content  string `desc:"内容"`
	}
)

func WithSVNPath(path string) SVNOption {
	return func(c *SVNConfig) {
		c.SVNPath = path
	}
}

func WithSVNUrl(url string) SVNOption {
	return func(c *SVNConfig) {
		c.URL = url
	}
}

func WithSVNUser(user string) SVNOption {
	return func(c *SVNConfig) {
		c.User = user
	}
}

func WithSVNPasswd(passwd string) SVNOption {
	return func(c *SVNConfig) {
		c.Passwd = passwd
	}
}
