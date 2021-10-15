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
	seelog.Info("request=>",ctx.Request.GetBody)
	response.Success(ctx,constant.GetSuccessCode,constant.GetSuccessMsg,gin.H{"message":"pong"})
}