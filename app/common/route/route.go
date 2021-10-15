package route

import (
	"firstgo/app/controller"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)


func PathRoute(r *gin.Engine) *gin.Engine{
	rootPath := r.Group("/gin")
	{
		pingPath := rootPath.Group("/ping")
		{
			fmt.Println(pingPath)
			seelog.Info(pingPath)
			controller.GetPing(pingPath)
		}
	}
	return r
}