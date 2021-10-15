package excel

import (
	"github.com/cihub/seelog"
	"github.com/xuri/excelize/v2"
)

func ReadExcel(filePath string){
	excelFile, err := excelize.OpenFile(filePath)
	if err != nil {
		seelog.Error(err)
	}
	var sheetMap map[int]string = excelFile.GetSheetMap()
	for sheetId,sheetName := range sheetMap{
		seelog.Info(sheetId,"=",sheetName)
	}

}