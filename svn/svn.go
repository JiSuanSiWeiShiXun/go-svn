package svn

import (
	"bytes"
	"io"
	"os/exec"
	"strings"

	"github.com/JiSuanSiWeiShiXun/log"
	"github.com/sirupsen/logrus"
)

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
	args := []string{"checkout", sc.URL}
	if localPath != nil {
		args = append(args, localPath[0])
	}
	args = append(args, sc.GetAuthOption()...)

	cmd := exec.Command(sc.SVNPath, args...)
	cmd.Stdout = log.Logger.WriterLevel(logrus.DebugLevel)
	cmd.Stderr = log.Logger.WriterLevel(logrus.ErrorLevel)
	log.Debug("[invoke] %s", cmd.String())
	return cmd.Run()
}

// SVNBlame 获取指定文件的每行最近提交信息
func (sc SVNClient) Blame(file string) (svnBlames []*SVNBlame, err error) {
	svnBlames = make([]*SVNBlame, 0)
	args := []string{"blame", "-v", file}
	args = append(args, sc.GetAuthOption()...)

	cmd := exec.Command(sc.SVNPath, args...)
	var buf bytes.Buffer
	writers := io.MultiWriter(log.Logger.WriterLevel(logrus.DebugLevel), &buf)
	cmd.Stdout = writers
	cmd.Stderr = log.Logger.WriterLevel(logrus.ErrorLevel)
	log.Debug("[invoke] %s", cmd.String())

	if err = cmd.Run(); err != nil {
		return
	}

	// 逐行解析
	for {
		var line string
		line, err = buf.ReadString('\t')
		if err != nil {
			if err == io.EOF {
				err = nil
				return
			}
			log.Error("read from buf error: %v", err)
			return
		}
		lineSlice := strings.Fields(line)
		svnBlame := &SVNBlame{
			Revision: lineSlice[0],
			Author:   lineSlice[1],
			Date:     strings.Join(lineSlice[2:5], " "), // e.g. `2022-11-24 14:03:08 +0800 (Thu, 24 Nov 2022)`
			Content:  strings.Join(lineSlice[6:], " "),
		}
		// fmt.Printf("[0]%v [1]%v [2]%v [3]%v\n", svnBlame.Revision, svnBlame.Author, svnBlame.Date, svnBlame.Content)
		svnBlames = append(svnBlames, svnBlame)
	}
}
