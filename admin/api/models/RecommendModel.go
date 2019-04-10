package models

import (
	"github.com/astaxie/beego/logs"
	"xstv/cloud/cms-admin/common/raw"
)

type RecommendModel struct {
	BaseModel
}

var (
	recommend_video_os = BaseModel{TABLE_NAME: "recommend_video"}
)

func (this *RecommendModel) RecommendAddVideo (bulkList []raw.RecommendVideo) {
	recommend_video_os._init()
	recommend_video_os.ORM_CON.Begin()
	_, err := recommend_video_os.ORM_CON.InsertMulti(1, bulkList)
	if err != nil {
		recommend_video_os.ORM_CON.Rollback()
		logs.Error(err)
	} else {
		recommend_video_os.ORM_CON.Commit()
	}
}