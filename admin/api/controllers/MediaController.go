package controllers

import (
	"strconv"
	"time"
	"xstv/cloud/cms-admin/admin/api/models"
	"xstv/cloud/cms-admin/common/raw"
)

type MediaController struct {
	BaseController
	models.MediaModel
}

type MediaListRows struct {
	List     []raw.MediaSearchVideo `json:"list"`
	Total    int64                  `json:"total"`
	Page     int                    `json:"page"`
	PageSize int                    `json:"pageSize"`
}

func (self *MediaController) _Init() {
	self._init()
}

func (self *MediaController) List() {
	defer self.Catch()
	title := self.GetString("title")
	year, _ := self.GetInt("year")
	director := self.GetString("director")
	cid, _ := self.GetInt("cid")
	status, _ := self.GetInt("status", -1)
	page, _ := self.GetInt("page", 1)
	pagesize, _ := self.GetInt("pagesize", 30)
	topic, _ := self.GetInt("topic")
	oparams := models.MediaOptions{
		Title:        title,
		Pubilsh_year: year,
		Directors:    director,
		CategoryId:   cid,
		Status:       status,
		Offset:       (page - 1) * pagesize,
		Limit:        pagesize,
		Topic:        topic,
	}
	count, Lists := self.MediaListInfo(oparams)
	MediaList := MediaListRows{
		Total:    count,
		Page:     page,
		PageSize: pagesize,
	}
	if len(Lists) == 0 {
		MediaList.List = []raw.MediaSearchVideo{}
		self.response.Data = MediaList
		self.Data["json"] = self.response
		self.ServeJSON()
	}
	for _, value := range Lists {
		createTime, _ := strconv.ParseInt(value.CreateTime, 10, 64)
		updateTime, _ := strconv.ParseInt(value.UpdateTime, 10, 64)
		item := raw.MediaSearchVideo{
			Id:           value.Id,
			Title:        value.Title,
			CategoryId:   value.CategoryId,
			CategoryName: value.CategoryName,
			Directors:    value.Directors,
			ReleaseYear:  value.ReleaseYear,
			Language:     value.Language,
			Status:       value.Status,
			CreateTime:   time.Unix(createTime, 0).Format("2006-01-02 15:04:05"),
			UpdateTime:   time.Unix(updateTime, 0).Format("2006-01-02 15:04:05"),
			Hide:         value.Hide,
		}
		MediaList.List = append(MediaList.List, item)
	}
	//MediaList.List = Lists
	self.response.Data = MediaList
	self.Data["json"] = self.response
	self.ServeJSON()
}
