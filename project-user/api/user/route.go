package user

import (
	"github.com/gin-gonic/gin"
	"github.com/xszhangxiacuo/PMS/project-user/router"
	"log"
)

func init() {
	log.Println("init user router")
	router.Register(&RouterUser{})
}

type RouterUser struct {
}

func (*RouterUser) Route(r *gin.Engine) {
	uh := NewUserHandler()
	r.POST("/project/login/getCaptcha", uh.getCaptcha)

}
