package main

import (
	"github.com/gin-gonic/gin"
	srv "github.com/xszhangxiacuo/PMS/project-common"
	_ "github.com/xszhangxiacuo/PMS/project-user/api"
	"github.com/xszhangxiacuo/PMS/project-user/router"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	srv.Run(r, "project-user", ":80")
}
