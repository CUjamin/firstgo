package controller

import (
	"firstgo/app/common/constant"
	"firstgo/app/common/response"

	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

type PingController struct{

}

func GetPing(pingGin *gin.RouterGroup){
	pingController :=&PingController{}
	pingGin.Use().GET("/ping",pingController.ping)
}

func (c PingController) ping(ctx *gin.Context){
	seelog.Info("request=>",ctx.Request)
	seelog.Info("request Mothed=>",ctx.Request.Method)
	seelog.Info("request Url=>",ctx.Request.URL.Path)
	seelog.Info("request Body=>",ctx.Request.GetBody)
	seelog.Info("request id=>",ctx.Query("id"))
	seelog.Info("request name=>",ctx.Query("name"))


	response.Success(ctx,constant.GetSuccessCode,constant.GetSuccessMsg,gin.H{"message":"pong"})
}