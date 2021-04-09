package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type ExternalImportData struct {
	Id            int       `orm:"column(cl_Id);pk"`
	Type          int       `orm:"column(cl_Type)"`
	Titles        string    `orm:"column(cl_Titles)"`
	JsonData      string    `orm:"column(cl_JsonData);type(json)"`
	CreateTime    time.Time `orm:"column(cl_CreateTime);auto_now_add;type(datetime)"`
	FieldCount    int       `orm:"column(cl_FieldCount)"`
	WhereTime     time.Time `orm:"column(cl_WhereTime);type(datetime)"`
	FirstIdField  string    `orm:"column(cl_FirstIdField)"`
	SecondIdField string    `orm:"column(cl_SecondIdField)"`
	Guid          string    `orm:"column(cl_Guid)"`
	ExtData1      string    `orm:"column(cl_ExtData1)"`
	ExtData2      string    `orm:"column(cl_ExtData2)"`
}

func (u *ExternalImportData) TableName() string {
	return "tab_external_import_data"
}
