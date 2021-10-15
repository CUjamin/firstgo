package main

import (
	"firstgo/app/common/route"
	"firstgo/app/excel"
	"os"
	"strings"

	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

func main(){
	initLog("seelog.xml")
	seelog.Info("start ...")
	defer seelog.Flush()
	initWeb()
	printExcel()
	seelog.Info("start ok")
}

func initLog(logConfigName string){
	var projectPath string= ""
	projectPath,_=os.Getwd()
	var logConfigPath string = strings.Join([]string{projectPath,"config",logConfigName},"/")
	logger, err:=seelog.LoggerFromConfigAsFile(logConfigPath)
	if err != nil{
		panic(err)
	}
	seelog.ReplaceLogger(logger)
	seelog.Info("init log OK")
}


func initWeb() {
	r := gin.New()
	r = route.PathRoute(r)
	r.Run()

}
func initWeb2() {
	seelog.Info("start web ...")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
	seelog.Info("start web ok")
}

func printExcel(){
	var filePath string = "test.xlsx"
	excel.ReadExcel(filePath)
}