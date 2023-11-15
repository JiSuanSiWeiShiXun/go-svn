package main

import (
	"github.com/JiSuanSiWeiShiXun/go-svn/svn"
)

func main() {
	opt := svn.SVNOption{
		SVNRepo: svn.SVNRepo{
			URL:    "https://xsjreposvr7.seasungame.com/svn/FSM/trunk/MechaWar/Client/Projects/Runtime/Settings/Table",
			User:   "wangjiaojiao1",
			Passwd: "wjflying0.ai",
		},
	}
	sc, err := svn.NewSVNClient(&opt)
	if err != nil {
		panic(err)
	}
	err = sc.Checkout(`./temp`)
	if err != nil {
		panic(err)
	}
}
