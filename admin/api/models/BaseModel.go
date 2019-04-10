package models

import (
	"xstv/cloud/cms-admin/common/raw"

	"github.com/astaxie/beego/orm"
)

type BaseModel struct {
	TABLE_NAME string
	DB_NAME    string
	ORM_CON    orm.Ormer
	OS         orm.QuerySeter
	COND       *orm.Condition
}

func (this *BaseModel) _init() {
	this.ORM_CON = orm.NewOrm()
	if this.DB_NAME != "" {
		dbErr := this.ORM_CON.Using(this.DB_NAME)
		if dbErr != nil {
			panic(dbErr)
		}
	}
	this.OS = this.ORM_CON.QueryTable(this.TABLE_NAME)
	this.COND = orm.NewCondition()
}

func (this *BaseModel) _InitMyqlConnect(table_name string) BaseModel {
	this.ORM_CON = orm.NewOrm()
	this.OS = this.ORM_CON.QueryTable(table_name)
	this.COND = orm.NewCondition()
	return *this
}

func init() {
	orm.RegisterModel(
		&raw.Category{},
		&raw.CategoryFiltertype{},
		&raw.CategoryFilteritem{},
		&raw.ThemeDesktop{}, //主题
		&raw.WrapDesktop{},  //专题
		&raw.Filters{},
		&raw.Topic{},
		&raw.PromotionDesktop{},
		&raw.LayoutTemplate{},
		&raw.TopicDesktop{},
		&raw.Media{},
		&raw.TopicContent{},
		&raw.TopicContentBlackList{},
		&raw.XsTagTop{},
		&raw.RecommendVideo{},
	)
}
