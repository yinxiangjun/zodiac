package models

import (
	"xstv/cloud/xsmedia/common/raw"
	"xstv/cloud/xsmedia/common/raw/std"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type BaseModel struct {
	TABLE_NAME string
	DB_NAME    string
	ORM_CON    orm.Ormer
	OS         orm.QuerySeter
	COND       *orm.Condition
}

var (
	ORM_ERR = map[error]string{
		orm.ErrNoRows:       "no row found",
		orm.ErrArgs:         "args error may be empty",
		orm.ErrMissPK:       "missed pk value",
		orm.ErrMultiRows:    "return multi rows",
		orm.ErrNotImplement: "have not implement",
		orm.ErrStmtClosed:   "stmt already closed",
		orm.ErrTxDone:       "transaction not begin",
		orm.ErrTxHasBegan:   "transaction already begin",
	}
)

func (this *BaseModel) _init() {
	this.ORM_CON = orm.NewOrm()
	if this.DB_NAME != "" {
		dbError := this.ORM_CON.Using(this.DB_NAME)
		if dbError != nil {
			panic(dbError)
		}
	}
	this.OS = this.ORM_CON.QueryTable(this.TABLE_NAME)
	this.COND = orm.NewCondition()
}

func (this *BaseModel) DeleteById(field string, id int) (num int64) {
	this._init()
	var err error
	num, err = this.OS.Filter(field, id).Delete()
	if err != nil {
		this.GovernErr(err)
	}
	return
}

func (this *BaseModel) Count() (int64, error) {
	return this.OS.Count()
}

func (this *BaseModel) GovernErr(e error) string {
	var errorString string = ORM_ERR[e]
	if errorString == "" {
		errorString = e.Error()
	}
	logs.Critical(errorString)
	logs.Error(errorString)
	return errorString
}

func init() {
	orm.RegisterModel(
		&raw.Topic{},
		&raw.Category{},
		&raw.CategoryAttr{},
		&raw.CleanStatus{},
		&raw.Video{},
		&raw.People{},
		&raw.Tag{},
		&raw.TopicVideo{},
		&raw.RecommendVideo{},
		&raw.Filters{},
		&raw.Category_v2{},        //分类
		&raw.Area{},               //地区
		&raw.Area_alias{},         //地区别名
		&raw.Language{},           //语言
		&raw.Language_alias{},     //语言别名
		&raw.Cp_v2{},              //cp
		&raw.Tag_v2{},             //类型标签
		&raw.Tag_alias{},          //类型标签别名
		&raw.Origin{},             //来换
		&raw.Poster{},             //海报
		&raw.Poster_address{},     //海报地址
		&raw.Media{},              //媒资
		&raw.SourceMediaV2{},      //媒资来源库
		&raw.Xsperson{},           //人物库
		&raw.Xsepisode{},          //子集库
		&raw.SourceEpisodeV2{},    //来源子集
		&raw.MediaPolymerV2{},     //待聚合节目
		&std.StdMedia{},           //第三方标准库
		&raw.Distribute_v2{},      //分发业务
		&raw.Distribute_data_v2{}, //分发数据表
		&raw.ImportSign{},         //批量导入查询
		&raw.AudioStation{},
		&raw.XsTag{},      //新视标签
		&raw.XtTagAlias{}, //新视标签同义词
		&raw.SeriesMap{},  //系列
		&raw.XsTagTop{},   //top榜

	)
}
