package controllers

import (
	"context"
	"fmt"
	"os"
	"reflect"

	"go-qlx-tool/models"
	iutils "go-qlx-tool/utils"

	"time"

	"github.com/beego/beego/v2/client/cache"
	"github.com/beego/beego/v2/client/orm"
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
	// picName := "/export/suningb2_2021_03_14_00_00_00-2021_03_14_00_59_59.xlsx"

	// fmt.Println(iutils.GetFileRootPath(picName))

	// client, err := oss.New("https://oss-cn-hangzhou.aliyuncs.com", "LTAI4GJdSz44ji9hGUxUT9RN", "03zgvRQ3zq2xj8uggPvOXsThVCcEhM")
	// if err != nil {
	// 	// HandleError(err)
	// 	fmt.Println(err)
	// }

	// lsRes, err := client.ListBuckets()
	// if err != nil {
	// 	// HandleError(err)
	// 	fmt.Println(err)
	// }

	// for _, bucket := range lsRes.Buckets {
	// 	fmt.Println("Buckets:", bucket.Name)
	// }

	// bucket, err := client.Bucket("daxun-emall")
	// if err != nil {
	// 	// HandleError(err)
	// 	fmt.Println(err)
	// }

	// // err = bucket.PutObject("my-object", "/export/suningb2_2021_03_14_00_00_00-2021_03_14_00_59_59.xlsx")
	// err = bucket.PutObjectFromFile("my-object1.xlsx", iutils.GetFileRootPath(picName))
	// if err != nil {
	// 	// HandleError(err)
	// 	fmt.Println(err)
	// }

	// c.Ctx.WriteString("200")
	// return

	o := orm.NewOrm()

	// model := models.ExternalImportData{Id: 2895489}

	// err := o.Read(&model)

	// if err == orm.ErrNoRows {
	// 	fmt.Println("查询不到")
	// } else if err == orm.ErrMissPK {
	// 	fmt.Println("找不到主键")
	// } else {
	// 	fmt.Println(model.Id, model.Titles)
	// }
	// c.Ctx.WriteString("200")
	// return

	// startTime := "2021-04-01 00:00:00"
	// startTime1, err := time.Parse("2006-01-02 15:04:05", startTime)
	// var model2 models.ExternalImportData
	// qs := o.QueryTable(new(models.ExternalImportData)) // 返回 QuerySeter
	// qs = qs.Filter("cl_FirstIdField", "DdbXT774")
	// qs = qs.Filter("cl_WhereTime__gte", "2021-04-01")
	// qs = qs.Filter("cl_WhereTime__lte", "2021-04-02")
	// count, err := qs.Count() // WHERE id = 1
	// if err == orm.ErrMultiRows {
	// 	// 多条的时候报错
	// 	fmt.Printf("Returned Multi Rows Not One")
	// }
	// if err == orm.ErrNoRows {
	// 	// 没有找到记录
	// 	fmt.Printf("Not row found")
	// }
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(err)
	// fmt.Println(count)
	// fmt.Println(model2)
	// c.Ctx.WriteString("200")
	// return

	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("*,(select cl_ActUserAccount from tab_activity_channel where cl_ActCode=tab_external_import_data.cl_FirstIdField and  cl_ChannelId=tab_external_import_data.cl_SecondIdField) cl_ActUserAccount").
		From("tab_external_import_data").
		Where("cl_FirstIdField = ? and cl_WhereTime between ? and ?").
		OrderBy("cl_Id").
		Limit(10).Offset(0)

	// 导出 SQL 语句
	sql := qb.String()
	fmt.Println(sql)

	var dataMaps []orm.Params

	// var extImportDatas []models.ExternalImportData
	// 执行 SQL 语句

	// ids := []interface{"DdbXT774", "2021-04-01", "2021-04-06"}

	params := []interface{}{"DdbXT774", "2021-04-01", "2021-04-06"}
	// params = append(params, 1)
	fmt.Println(params)

	num, err := o.Raw(sql, params).Values(&dataMaps)

	fmt.Println(err)
	fmt.Println(num)
	fmt.Println(dataMaps[0])
	c.Ctx.WriteString("200")
	return

	// var maps []orm.Params
	// num, _ := o.Raw("SELECT * FROM user").Values(&maps)
	// for _, term := range maps {
	// 	fmt.Println(term["id"], ":", term["name"])

	// 	c.Ctx.WriteString("200")
	// 	return
	// }
	// fmt.Println(num)
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

func (c *ExportController) ExportExternalImportData() {
	if m_bm == nil {
		m_bm, _ = cache.NewCache("memory", `{"interval":10}`)
	}

	// LOCK, _ := m_bm.Get(context.TODO(), "ExportExternalImportData_LOCK")
	// if LOCK == nil {
	// 	m_bm.Put(context.TODO(), "ExportExternalImportData_LOCK", 1, 20*time.Second)
	// } else {
	// 	c.Redirect("/error?msg=重复执行请稍后...&returl=/qulaxin", 302)
	// 	// c.Ctx.WriteString("重复执行请稍后...")
	// 	return
	// }

	firstIdField := c.GetString("firstid")
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

	var exportTitles = [20]string{"活动名称", "活动代码", "推客帐号", "关系备注", "日期", "标识", "数据1", "数据2", "数据3", "数据4", "数据5", "应付款", "导入结算"}

	for i := 0; i < len(exportTitles); i++ {
		if exportTitles[i] == "" {
			break
		}
		cell = row.AddCell()
		cell.Value = exportTitles[i]
	}

	forNum := 0
	for {
		fmt.Println(pageId + cast.ToString(forNum))

		o := orm.NewOrm()

		var dataMaps []orm.Params

		var modelList []models.ExternalImportData
		qs := o.QueryTable(new(models.ExternalImportData)) // 返回 QuerySeter
		if firstIdField != "" {
			qs = qs.Filter("cl_FirstIdField", firstIdField)
		}
		qs = qs.Filter("cl_WhereTime__gte", startTime)
		qs = qs.Filter("cl_WhereTime__lte", endTime)

		num, err := qs.Values(&dataMaps)
		// num, err := qs.All(&modelList)
		if err == orm.ErrMultiRows {
			// 多条的时候报错
			fmt.Printf("Returned Multi Rows Not One")
		}
		if err == orm.ErrNoRows {
			// 没有找到记录
			fmt.Printf("Not row found")
		}
		if err != nil {
			fmt.Println(err)
		}

		bT := time.Now() // 开始时间

		var actCodes []string
		var dataMaps2 []orm.Params
		for _, value := range dataMaps {
			// fmt.Print(index, "\t")
			// fmt.Print(value, "\t")

			// actCodes = append(actCodes, value["FirstIdField"].(string))
			actCodes = iutils.AddArrayEx(actCodes, value["FirstIdField"].(string))
			dataMaps2 = append(dataMaps2, value)
		}

		eT := time.Since(bT) // 从开始到当前所消耗的时间

		fmt.Println("Run time: ", eT)
		fmt.Println(len(actCodes))

		var actChannelMaps []orm.Params

		qsActChannel := o.QueryTable(new(models.ActivityChannel))
		qsActChannel.Filter("cl_ActCode__in", actCodes).Values(&actChannelMaps)

		fmt.Println(len(actChannelMaps))

		actChannelMap2 := make(map[string]interface{})
		for _, value := range actChannelMaps {
			fmt.Println(value)
			actChannelMap2[fmt.Sprintf("%v_%v", value["ActCode"], value["ChannelId"])] = value
			fmt.Println(actChannelMap2)
			break
		}

		fmt.Println(len(actChannelMaps))
		fmt.Println(len(dataMaps2))
		c.Ctx.WriteString("200")

		return

		fmt.Println(num)
		fmt.Println(len(modelList))

		// value, ok := resultMap["data"]
		// if !ok {
		// 	c.Ctx.WriteString("数据错误：" + resultStr)
		// 	return
		// }

		// dataMap := cast.ToStringMap(value)

		// pageId = cast.ToString(dataMap["page_id"])

		// if forNum == 0 {
		// 	value, _ = dataMap["export_title"]

		// 	exportTitle := reflect.ValueOf(value)

		// 	for i := 0; i < exportTitle.Len(); i++ {
		// 		cell = row.AddCell()
		// 		cell.Value = exportTitle.Index(i).Elem().String()
		// 	}
		// }

		// value, _ = dataMap["export_data"]
		// exportData := reflect.ValueOf(value)

		// if exportData.Len() < 1 {
		// 	break
		// }
		// fmt.Println(exportData.Len())

		for i := 0; i < len(modelList); i++ {
			row = sheet.AddRow()

			model := modelList[i]
			cell = row.AddCell()
			cell.Value = "huodong"

			cell = row.AddCell()
			cell.Value = model.FirstIdField

			cell = row.AddCell()
			cell.Value = ""

			break
		}
		forNum++
		break
	}

	path := "export/external_import_data"
	if !utils.FileExists(path) {
		os.MkdirAll(path, os.ModePerm)
	}

	startT, _ := time.Parse("2006-01-02 15:04:05", startTime)
	endT, _ := time.Parse("2006-01-02 15:04:05", endTime)
	filename := path + "external_import_data_" + startT.Format("2006_01_02_15_04_05") + "-" + endT.Format("2006_01_02_15_04_05") + ".xlsx"

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
