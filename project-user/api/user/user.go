package user

import (
	"context"
	"github.com/gin-gonic/gin"
	common "github.com/xszhangxiacuo/PMS/project-common"
	"github.com/xszhangxiacuo/PMS/project-user/internal/dao"
	"github.com/xszhangxiacuo/PMS/project-user/internal/repo"
	"github.com/xszhangxiacuo/PMS/project-user/pkg/model"
	"log"
	"net/http"
	"strconv"
	"time"
)

type UserHandler struct {
	cache repo.Cache
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		cache: dao.Rc,
	}
}

// 获取验证码
func (uh *UserHandler) getCaptcha(ctx *gin.Context) {
	rsp := &common.Result{}
	//1.获取参数
	mobile := ctx.PostForm("mobile")
	//2.校验参数
	if !common.VerifyMobile(mobile) {
		ctx.JSON(http.StatusOK, rsp.Fail(model.NoLegalMobile, "手机号不合法"))
	}
	//3.生成验证码（随机4位1000-9999）
	code := common.RandomNum(1000, 9999)
	//4.调用短信平台（用go协程执行，接口可以快速响应）
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("短信发送成功")
		//5.将验证码存入redis，验证码有效时间15min
		c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err := uh.cache.Put(c, "REGISTER_"+mobile, strconv.Itoa(code), 15*time.Minute)
		if err != nil {
			log.Printf("验证码存入redis出错，cause by:%v\n", err)
		}
		log.Printf("手机号和验证码成功存入redis：REGISTER_%s:%s", mobile, strconv.Itoa(code))
	}()
	//由于并没有调用短信发送平台，所以在这里要将验证码直接返回前端用于验证登录
	ctx.JSON(http.StatusOK, rsp.Success(code))
}
