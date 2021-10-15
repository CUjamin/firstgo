package main

import (
	"firstgo/app/excel"
	"os"
	"strings"
	"github.com/cihub/seelog"
)

func main(){
	initLog("seelog.xml")
	seelog.Info("start ...")
	defer seelog.Flush()

	var filePath string = "test.xlsx"
	excel.ReadExcel(filePath)
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