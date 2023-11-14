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

// SVNCheckout 检出svn指定路径的文件
func (sc SVNClient) Checkout(localPath ...string) error {
	cmdStr := []string{"checkout", sc.URL}
	if localPath != nil {
		cmdStr = append(cmdStr, localPath[0])
	}
	cmd := exec.Command(sc.SVNPath, cmdStr...)
	cmd.Stdout = log.Logger.Writer()
	cmd.Stderr = log.Logger.Writer()
	log.Debug(cmd.String())
	return cmd.Run()
}

func (sc SVNClient) Blame(file string) (map[string]string, error) {
	return nil, nil
}
