package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type ActivityChannel struct {
	Id             int       `orm:"column(cl_Id);pk"`
	ChannelId      string    `orm:"column(cl_ChannelId)"`
	ActCode        string    `orm:"column(cl_ActCode)"`
	ActUserAccount string    `orm:"column(cl_ActUserAccount)"`
	Remark         string    `orm:"column(cl_Remark)"`
	Link           string    `orm:"column(cl_Link)"`
	Status         int       `orm:"column(cl_Status)"`
	Source         int       `orm:"column(cl_Source)"`
	CreateTime     time.Time `orm:"column(cl_CreateTime);auto_now_add;type(datetime)"`
	UpdateTime     time.Time `orm:"column(cl_UpdateTime);auto_now;type(datetime)"`
	UseTime        time.Time `orm:"column(cl_UseTime);type(datetime)"`
}

func (u *ActivityChannel) TableName() string {
	return "tab_activity_channel"
}
