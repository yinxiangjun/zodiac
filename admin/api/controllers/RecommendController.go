package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strconv"
	"xstv/cloud/cms-admin/admin/api/models"
	"xstv/cloud/cms-admin/common/raw"
)

type RecommendController struct {
	BaseController
	models.RecommendModel
}

type RecommendListItem struct {
	ListId       int    `json:"list_id"`
	Name         string `json:"name"`
	Title        string `json:"title"`
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	Desc         string `json:"desc"`
}

type RecommendListData struct {
	List []RecommendListItem `json:"list"`
}

type RecommendVideoItem struct {
	Id             int64  `json:"id"`
	Title          string `json:"title"`
	RecommendOrder int    `json:"recommend_order"`
	GlobalOrder    int    `json:"global_order"`
	CostType       int    `json:"cost_type"`
	ImgUrl         string `json:"img_url"`
	Status         int    `json:"status"`
	Persons        string `json:"persons"`
	ReleaseYear    int    `json:"release_year"`
	CategoryName   string `json:"category_name"`
}

type RecommendDetailData struct {
	Total    int                  `json:"total"`
	List     []RecommendVideoItem `json:"list"`
	Page     int                  `json:"page"`
	PageSize int                  `json:"pageSize"`
}

type RecommendVideoRaw struct {
	Id           int
	ListId       int
	VideoId      int64
	StatusInList int
	OrderInList  int
	OriginScore  int
	Title        string
	OriginPoster string
	IsVip        int
	IsTvod       int
	CategoryName string
	Directors    string
	MainActors   string
	ReleaseYear  string
}

type AddVideoBody struct {
	List []struct {
		VideoId    int64  `json:"video_id"`
		VideoOrder int    `json:"video_order"`
		CostumImg  string `json:"custom_img"`
	} `json:"list"`
}

type RelVideo struct {
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	ImgUrl   string `json:"img_url"`
	CostType int    `json:"cost_type"`
}

type RelVideoList struct {
	List []RelVideo `json:"list"`
}

func (self *RecommendController) Lists() {
	self._init()
	defer self.Catch()

	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").From("recommend_list").OrderBy("category_id, name")
	sql := qb.String()
	var listData RecommendListData
	var lists []raw.RecommendList
	_, err := o.Raw(sql).QueryRows(&lists)
	if err != nil {
		panic(err)
	}
	if len(lists) == 0 {
		listData.List = []RecommendListItem{}
		self.response.Data = listData
		self.Data["json"] = self.response
		self.ServeJSON()
	}
	for _, list := range lists {
		listData.List = append(listData.List, RecommendListItem{
			ListId:       list.Id,
			Name:         list.Name,
			Title:        list.Title,
			CategoryId:   list.CategoryId,
			CategoryName: list.CategoryName,
			Desc:         list.Desc,
		})
	}
	self.response.Data = listData
	self.Data["json"] = self.response
	self.ServeJSON()
}

func (self *RecommendController) Detail() {
	self._init()
	defer self.Catch()

	listId, err := self.GetInt("list_id")
	if err != nil {
		panic("list_id:数据格式有误！")
	}
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	n, err := self.GetInt("pageSize")
	if err != nil {
		n = 20
	}

	o := orm.NewOrm()
	qbTotal, _ := orm.NewQueryBuilder("mysql")
	qbTotal.Select("count(*)").
		From("recommend_video as a").
		InnerJoin("xsmedia as b").On("a.video_id = b.id").
		Where("b.status=1 and a.list_id = ? and status_in_list=1").
		OrderBy("-status_in_list", "-order_in_list", "-b.origin_score")
	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("a.*, b.title, b.origin_poster, c.is_tvod, c.is_vip, "+
		"b.category_name, b.directors, b.main_actors, LEFT(b.release_time,4) as release_time, b.origin_score").
		From("recommend_video as a").
		InnerJoin("xsmedia as b").On("a.video_id = b.id").
		LeftJoin("source_media_v2 AS c").On("b.id=c.xs_media_id").
		Where("b.status=1 and a.list_id = ? and status_in_list=1").
		OrderBy("-status_in_list", "-order_in_list", "-b.origin_score").
		Limit(n).Offset(n * (page - 1))
	sqlTotal := qbTotal.String()
	var totalResult []int
	_, err = o.Raw(sqlTotal, listId).QueryRows(&totalResult)
	logs.Debug("===", totalResult)

	sql := qb.String()
	var detailData RecommendDetailData
	var videos []RecommendVideoRaw
	detailData.Total = totalResult[0]
	detailData.Page = page
	detailData.PageSize = n
	_, err = o.Raw(sql, listId).QueryRows(&videos)
	if err != nil {
		panic(err)
	}
	if len(videos) == 0 {
		detailData.List = []RecommendVideoItem{}
		self.response.Data = detailData
		self.Data["json"] = self.response
		self.ServeJSON()
	}
	for _, video := range videos {
		item := RecommendVideoItem{
			Id:             video.VideoId,
			Title:          video.Title,
			RecommendOrder: video.OrderInList,
			GlobalOrder:    video.OriginScore,
			ImgUrl:         video.OriginPoster,
			Status:         video.StatusInList,
			CategoryName:   video.CategoryName,
		}
		if video.IsTvod == 1 {
			item.CostType = 2
		} else if video.IsVip == 1 {
			item.CostType = 1
		} else {
			item.CostType = 0
		}
		item.Persons = video.Directors
		item.ReleaseYear, _ = strconv.Atoi(video.ReleaseYear)
		detailData.List = append(detailData.List, item)
	}
	self.response.Data = detailData
	self.Data["json"] = self.response
	self.ServeJSON()
}

func (self *RecommendController) Add() {
	self._init()
	defer self.Catch()

	listId, err := self.GetInt("list_id")
	if err != nil {
		panic("list_id:格式有误！")
	}
	var addL AddVideoBody
	bodyContent := self.Ctx.Input.RequestBody
	err = json.Unmarshal(bodyContent, &addL)
	if err != nil {
		panic("json格式有误！")
	}
	var bulkList []raw.RecommendVideo
	for _, item := range addL.List {
		bulkList = append(bulkList, raw.RecommendVideo{
			ListId:       listId,
			VideoId:      item.VideoId,
			StatusInList: 1,
			OrderInList:  item.VideoOrder,
			CustomImg:    item.CostumImg,
		})
	}
	self.RecommendAddVideo(bulkList)
	self.Data["json"] = self.response
	self.ServeJSON()
}

func (self *RecommendController) Remove() {
	self._init()
	defer self.Catch()

	listId, err := self.GetInt("list_id")
	if err != nil {
		panic("list_id:格式有误！")
	}
	videoId, err := self.GetInt64("video_id")
	if err != nil {
		panic("video_id:格式有误！")
	}
	o := orm.NewOrm()
	_, err = o.QueryTable("recommend_video").Filter("list_id", listId).Filter("video_id", videoId).Delete()
	if err != nil {
		logs.Error(err)
	}
	self.Data["json"] = self.response
	self.ServeJSON()
}

func (self *RecommendController) VideoReorder() {
	self._init()
	defer self.Catch()

	listId, err := self.GetInt("list_id")
	if err != nil {
		panic("list_id:格式有误！")
	}
	itemId, err := self.GetInt64("video_id")
	if err != nil {
		panic("video_id:格式有误！")
	}
	newOrder, err := self.GetInt("new_order")
	if err != nil {
		panic("new_order:格式有误！")
	}
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Update("recommend_video").Set("order_in_list = ?").Where("list_id = ? and video_id = ?")
	sql := qb.String()
	_, err = o.Raw(sql, newOrder, listId, itemId).Exec()
	if err != nil {
		logs.Error(err)
	}
	self.Data["json"] = self.response
	self.ServeJSON()
}
