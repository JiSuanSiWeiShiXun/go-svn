package svn

import (
	"os/exec"

	"github.com/JiSuanSiWeiShiXun/log"
)

type SVNRepo struct {
	URL    string
	User   string
	Passwd string
}

type SVNOption struct {
	SVNPath string // svn path
	SVNRepo
}

type SVNClient struct {
	SVNOption
}

// NewSVNClient 初始化svn操作对象
func NewSVNClient(opt *SVNOption) (*SVNClient, error) {
	if opt.SVNPath == "" {
		opt.SVNPath = "svn"
	}
	if opt.URL == "" {
		return nil, ErrInvalidURL
	}
	return &SVNClient{SVNOption: *opt}, nil
}

// GetAuthOption 返回鉴权的命令行参数
func (sc SVNClient) GetAuthOption() (options []string) {
    options = make([]string, 0)
    if sc.User != "" {
        options = append(options, "--username", sc.User)
    }
    if sc.Passwd != "" {
        options = append(options, "--password", sc.Passwd)
    }
    return
}

// SVNCheckout 检出svn指定路径的文件
func (sc SVNClient) Checkout(localPath ...string) error {
	cmdSlice := []string{"checkout", sc.URL}
	if localPath != nil {
	    cmdSlice = append(cmdSlice, localPath[0])
	}
	cmdSlice = append(cmdSlice, sc.GetAuthOption()...)

	cmd := exec.Command(sc.SVNPath, cmdSlice...)
	cmd.Stdout = log.Logger.Writer()
	cmd.Stderr = log.Logger.Writer()
	log.Debug(cmd.String())
	return cmd.Run()
}

func (sc SVNClient) Blame(file string) (map[string]string, error) {
	return nil, nil
}
