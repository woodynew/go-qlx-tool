package controllers

import (
	"context"
	"fmt"
	"os"
	"reflect"

	iutils "go-qlx-tool/utils"

	"time"

	"github.com/beego/beego/v2/client/cache"
	beego "github.com/beego/beego/v2/server/web"

	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/utils"
	"github.com/spf13/cast"
	"github.com/tealeg/xlsx"
)

var m_bm cache.Cache

// ExportController operations for Export
type ExportController struct {
	beego.Controller
}

func (c *ExportController) ExportTest() {
	c.Ctx.WriteString("200")
	return
}

func (c *ExportController) GetExportSuningB2() {
	if m_bm == nil {
		m_bm, _ = cache.NewCache("memory", `{"interval":60}`)
	}

	//TODO 进度条
	astaxie1, _ := m_bm.Get(context.TODO(), "astaxie")
	m_bm.Put(context.TODO(), "astaxie", 1, 10*time.Second)
	// bm.IsExist(context.TODO(), "astaxie")
	// bm.Delete(context.TODO(), "astaxie")

	fmt.Println(astaxie1)

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "export/suningb2.tpl"
}

func (c *ExportController) ExportSuningB2() {
	if m_bm == nil {
		m_bm, _ = cache.NewCache("memory", `{"interval":60}`)
	}

	ExportSuningB2_LOCK, _ := m_bm.Get(context.TODO(), "ExportSuningB2_LOCK")
	if ExportSuningB2_LOCK == nil {
		m_bm.Put(context.TODO(), "ExportSuningB2_LOCK", 1, 60*time.Second)
	} else {
		c.Redirect("/error?msg=重复执行请稍后...&returl=/qulaxin", 302)
		// c.Ctx.WriteString("重复执行请稍后...")
		return
	}

	// configEnv, _ := beego.AppConfig.String("runmode")
	configQlxApiUrl, _ := beego.AppConfig.String("qulaxinapihost")
	if configQlxApiUrl == "" {
		c.Ctx.WriteString("缺少配置：qulaxinapihost")
		return
	}

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
		fmt.Println(pageId + cast.ToString(forNum))

		apiUrl := configQlxApiUrl + "/export-tool-api/businessnewbuyerQuery"

		resultStr := ""
		resultMap := make(map[string]interface{})

		reqApiCount := 0
		for {
			if reqApiCount > 3 {
				c.Ctx.WriteString("接口最大次数请求错误")
				return
			}
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

			resultMap = iutils.JSONToMap(str)
			if nil == resultMap {
				c.Ctx.WriteString("数据解析错误：" + str)
				return
			}

			value, ok := resultMap["ret"]
			if !ok {
				c.Ctx.WriteString("响应错误：" + str)
				return
			}
			if !reflect.ValueOf(value).Bool() {
				value, _ := resultMap["error"]
				if cast.ToString(value) == "1099" {
					reqApiCount++
					time.Sleep(time.Second)
					continue
				}
				c.Ctx.WriteString("响应错误：" + str)
				return
			}
			resultStr = str
			break
		}

		if resultStr == "" {
			c.Ctx.WriteString("无数据")
			return
		}

		value, ok := resultMap["data"]
		if !ok {
			c.Ctx.WriteString("数据错误：" + resultStr)
			return
		}

		dataMap := cast.ToStringMap(value)

		pageId = cast.ToString(dataMap["page_id"])

		if forNum == 0 {
			value, _ = dataMap["export_title"]

			exportTitle := reflect.ValueOf(value)

			for i := 0; i < exportTitle.Len(); i++ {
				cell = row.AddCell()
				cell.Value = exportTitle.Index(i).Elem().String()
			}
		}

		value, _ = dataMap["export_data"]
		exportData := reflect.ValueOf(value)

		if exportData.Len() < 1 {
			break
		}
		fmt.Println(exportData.Len())

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

	startT, _ := time.Parse("2006-01-02 15:04:05", startTime)
	endT, _ := time.Parse("2006-01-02 15:04:05", endTime)
	filename := path + "suningb2_" + startT.Format("2006_01_02_15_04_05") + "-" + endT.Format("2006_01_02_15_04_05") + ".xlsx"

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
