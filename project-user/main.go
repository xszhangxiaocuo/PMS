package main

import (
	"github.com/gin-gonic/gin"
	srv "github.com/xszhangxiacuo/PMS/project-common"
	"github.com/xszhangxiacuo/PMS/project-common/logs"
	_ "github.com/xszhangxiacuo/PMS/project-user/api"
	"github.com/xszhangxiacuo/PMS/project-user/router"
	"log"
)

func main() {
	r := gin.Default()

	//从配置中读取日志配置，初始化日志
	lc := &logs.LogConfig{
		DebugFileName: "F:\\workSpace\\PMS\\logs\\debug\\project-debug.log",
		InfoFileName:  "F:\\workSpace\\PMS\\logs\\info\\project-info.log",
		WarnFileName:  "F:\\workSpace\\PMS\\logs\\error\\project-error.log",
		MaxSize:       500,
		MaxAge:        28,
		MaxBackups:    3,
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}

	router.InitRouter(r)
	srv.Run(r, "project-user", ":80")
}
