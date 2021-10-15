package response

import (
	"firstgo/app/common/constant"
	"net/http"
	"github.com/gin-gonic/gin"
)



type ResultVO struct{
	Code 		constant.ResponseCode 	`json:"code"`
	Msg 		constant.ResponseMsg 	`json:"msg"`
	Success 	bool 					`json:"success"`
	Data 		interface{} 			`json:"data"`
}

func Success(ctx *gin.Context,code constant.ResponseCode,msg constant.ResponseMsg,data interface{}){
	response:=&ResultVO{
		Code: code,
		Msg: msg,
		Success: true,
		Data: data,
	}
	ctx.JSON(http.StatusOK,response)
}

func Failure(ctx *gin.Context,code constant.ResponseCode,msg constant.ResponseMsg,data interface{}){
	response:=&ResultVO{
		Code: code,
		Msg: msg,
		Success: false,
		Data: data,
	}
	ctx.JSON(http.StatusInternalServerError,response)
}