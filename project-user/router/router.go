package router

import "github.com/gin-gonic/gin"

// Router 路由接口，实现Route方法完成路由注册
type Router interface {
	Route(r *gin.Engine)
}

// RegisterRouter 路由注册
type RegisterRouter struct {
}

func NewRegisterRouter() *RegisterRouter {
	return &RegisterRouter{}
}

func (*RegisterRouter) Route(ro Router, r *gin.Engine) {
	ro.Route(r)
}

var routers []Router

// InitRouter 统一注册路由组中的路由
func InitRouter(r *gin.Engine) {
	for _, ro := range routers {
		ro.Route(r)
	}
}

// Register 传入任意多个路由
func Register(ro ...Router) {
	routers = append(routers, ro...)
}
