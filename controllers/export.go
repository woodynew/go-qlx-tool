package controllers

import (
	"fmt"
	"os"
	"reflect"

	iutils "go-qlx-tool/utils"

	beego "github.com/beego/beego/v2/server/web"

	"time"

	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/utils"
	"github.com/spf13/cast"
	"github.com/tealeg/xlsx"
)

// ExportController operations for Export
type ExportController struct {
	beego.Controller
}

func (c *ExportController) ExportTest() {
	c.Ctx.WriteString("200")
	return
}
func (c *ExportController) ExportSuningB2() {
	startTime := c.GetString("start_time")
	endTime := c.GetString("end_time")
	if startTime == "" || endTime == "" {
		c.Ctx.WriteString("时间范围错误")
		return
	}

	pageId := ""
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	file = xlsx.NewFile()
	sheet, _ = file.AddSheet("sheet1")
	row = sheet.AddRow()

	forNum := 0
	for {
		fmt.Println(pageId)

		apiUrl := "http://172.28.165.193:10181/export-tool-api/businessnewbuyerQuery"
		// apiUrl := "http://localhost:9502/export-tool-api/businessnewbuyerQuery"
		req := httplib.Get(apiUrl)
		req.Param("start_time", startTime)
		req.Param("end_time", endTime)
		req.Param("page_id", pageId)
		// req.Debug(true)
		// req.Response()

		str, err := req.String()
		if err != nil {
			c.Ctx.WriteString("ERROR")
			return
		}

		resultMap := iutils.JSONToMap(str)
		if nil == resultMap {
			c.Ctx.WriteString("数据解析错误：" + str)
			return
		}

		value, ok := resultMap["error"]
		if !ok {
			c.Ctx.WriteString("响应错误：" + str)
			return
		}
		value, ok = resultMap["data"]
		if !ok {
			c.Ctx.WriteString("数据错误：" + str)
			return
		}

		dataMap := cast.ToStringMap(value)

		pageId = cast.ToString(dataMap["page_id"])

		if forNum == 0 {
			value, ok = dataMap["export_title"]

			exportTitle := reflect.ValueOf(value)

			for i := 0; i < exportTitle.Len(); i++ {
				cell = row.AddCell()
				cell.Value = exportTitle.Index(i).Elem().String()
			}
		}

		value, ok = dataMap["export_data"]
		exportData := reflect.ValueOf(value)

		if exportData.Len() < 1 {
			break
		}

		for i := 0; i < exportData.Len(); i++ {
			row = sheet.AddRow()

			tempData := reflect.ValueOf(exportData.Index(i).Interface())

			for i2 := 0; i2 < tempData.Len(); i2++ {
				cell = row.AddCell()

				// tempValue := reflect.ValueOf(tempData.Index(i2).Interface())
				tempValue := tempData.Index(i2)
				// fmt.Println(reflect.ValueOf(tempValue.Interface()).Type())
				if reflect.ValueOf(tempValue.Interface()).Type().Kind() == reflect.Float64 {
					cell.Value = iutils.BigFloatToString(reflect.ValueOf(tempValue.Interface()).Float())
				} else {
					cell.Value = reflect.ValueOf(tempValue.Interface()).String()
				}
			}

		}
		forNum++
	}

	path := "export/"
	if !utils.FileExists(path) {
		os.MkdirAll(path, os.ModePerm)
	}
	filename := path + cast.ToString(time.Now().Unix()) + ".xlsx"
	file.Save(filename)

	c.Ctx.Output.Download(filename)
	c.Ctx.WriteString("200")
}

func ExportExcel() (filename string, err error) {

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	file = xlsx.NewFile()
	sheet, _ = file.AddSheet("sheet1")
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "第一行第一列"

	cell = row.AddCell()
	cell.Value = "第一行第二列"

	for i := 0; i < 100; i++ {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = "第二行第一列"
		cell = row.AddCell()
		cell.Value = "第二行第二列"
	}

	path := "export/"
	if !utils.FileExists(path) {
		os.MkdirAll(path, os.ModePerm)
	}
	filename = path + cast.ToString(time.Now().Unix()) + ".xlsx"
	err = file.Save(filename)
	return filename, err
}
