package main

import (
	"github.com/JiSuanSiWeiShiXun/go-svn/svn"

	"github.com/JiSuanSiWeiShiXun/log"
)

func main() {
	sc, err := svn.SVNFactory{}.NewClient(
		svn.WithSVNUrl("https://xsjreposvr7.seasungame.com/svn/FSM/trunk/MechaWar/Client/Projects/Runtime/Settings/Table"),
		svn.WithSVNUser("wangjiaojiao1"),
		svn.WithSVNPasswd("wjflying0.ai"),
	)
	if err != nil {
		panic(err)
	}
	if err = sc.Checkout(`./temp`); err != nil {
		panic(err)
	}

	svnBlames, err := sc.Blame(`./temp/GachaControl.tab`)
	if err != nil {
		panic(err)
	}

	cnt := 2
	log.Info("[line]%d 最新提交信息 %+v", cnt, svnBlames[cnt])
}
