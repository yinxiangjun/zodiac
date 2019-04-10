package models

import (
	"cloud/common/raw"
	"strconv"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type MediaModel struct {
	BaseModel
}

var (
	media_os = BaseModel{TABLE_NAME: new(raw.Media).TableName()}
)

type MediaOptions struct {
	Id              int
	Ids             []int
	Title           string
	CategoryId      int
	Directors       string
	Status          int
	Type            string
	Pubilsh_year    int
	CreateStartTime int
	CreateEndTime   int
	UpdateStartTime int
	UpdateEndTime   int
	Offset          int
	Limit           int
}

// type MediaAttr struct {
// 	Id         int64
// 	Title       string
// 	Category_name      string
// 	Directors       string
// 	Heat     int
// 	Total_episodes int
// 	Status int
// 	Create_time int
// }

func (this *MediaModel) ReadMediaInfo(info MediaOptions) (int64, []raw.Media) {
	media_os._init()
	if info.Title != "" {
		return this.SpeedSearchByPosition(info)
	}

	if info.Id > 0 {
		media_os.OS = media_os.OS.Filter("id", info.Id)
	}
	if len(info.Ids) > 0 {
		media_os.OS = media_os.OS.Filter("id__in", info.Ids)
	}

	if info.CategoryId > 0 {
		media_os.OS = media_os.OS.Filter("category_id", info.CategoryId)
	}
	if info.Directors != "" {
		media_os.OS = media_os.OS.Filter("directors__contains", info.Directors)
	}
	if info.Pubilsh_year > 0 {
		media_os.OS = media_os.OS.Filter("publish_year", info.Pubilsh_year)
	}
	if info.Type != "" {
		media_os.OS = media_os.OS.Filter("type__contains", info.Type)
	}
	if info.Status > 0 {
		media_os.OS = media_os.OS.Filter("status", info.Status)
	}

	if info.CreateStartTime > 0 && info.CreateEndTime > 0 {
		media_os.OS = media_os.OS.Filter("create_time__gte", info.CreateStartTime).Filter("create_time__lte", info.CreateEndTime)
	}
	if info.UpdateStartTime > 0 && info.UpdateEndTime > 0 {
		media_os.OS = media_os.OS.Filter("update_time__gte", info.UpdateStartTime).Filter("update_time__lte", info.UpdateEndTime)
	}

	if info.Limit <= 0 {
		info.Limit = 20
	}

	templateList := []raw.Media{}
	media_os.OS = media_os.OS.OrderBy("-id").Offset(info.Offset).Limit(info.Limit)
	_, error := media_os.OS.All(&templateList)

	if error != nil {
		logs.Error(error)
		return 0, templateList
	}
	count, err := media_os.OS.Count()
	if err != nil {
		logs.Error(err)
	}
	return count, templateList
}

func (this *MediaModel) SpeedSearchByPosition(info MediaOptions) (int64, []raw.Media) {
	media_os._init()
	valueSlice := []interface{}{}
	templateList := []raw.Media{}
	prefix := "SELECT id,title FROM " + media_os.TABLE_NAME
	where := " WHERE 1=1 "
	if info.Title != "" {
		where += "AND position(? in `title`)  "
		valueSlice = append(valueSlice, info.Title)
	}
	if info.CategoryId > 0 {
		where += "AND category_id = ? "
		valueSlice = append(valueSlice, info.CategoryId)
	}

	if info.Directors != "" {
		where += "AND directors like '%" + info.Directors + "%'"
		//where += "AND directors like \"%?%\"   "
		//valueSlice = append(valueSlice, info.Directors)
	}

	if info.Pubilsh_year > 0 {
		where += "AND publish_year = ? "
		valueSlice = append(valueSlice, info.Pubilsh_year)
	}

	if info.Type != "" {
		where += "AND type like \"%?%\" "
		valueSlice = append(valueSlice, info.Type)
	}

	if info.Status > 0 {
		where += "AND status = ? "
		valueSlice = append(valueSlice, info.Status)
	}

	if info.CreateStartTime > 0 && info.CreateEndTime > 0 {
		where += "AND create_time >= ? AND create_time <= ? "
		valueSlice = append(valueSlice, info.CreateStartTime, info.CreateEndTime)
	}

	if info.UpdateStartTime > 0 && info.UpdateEndTime > 0 {
		where += "AND update_time >= ? AND update_time <= ? "
		valueSlice = append(valueSlice, info.UpdateStartTime, info.UpdateEndTime)
	}

	if info.Limit <= 0 {
		info.Offset = 0
		info.Limit = 20
	}
	limit := "LIMIT ?,? "
	orderby := "ORDER BY id DESC "
	sql := prefix + where + orderby + limit
	values := append(valueSlice, info.Offset, info.Limit)
	_, error := media_os.ORM_CON.Raw(sql, values...).QueryRows(&templateList)

	if error != nil {
		logs.Error(error)
		return 0, templateList
	}
	if len(templateList) <= 0 {
		return 0, templateList
	}

	ids := []int64{}
	for _, value := range templateList {
		ids = append(ids, value.Id)
	}

	result := []raw.Media{}
	_, queryErr := media_os.OS.Filter("id__in", ids).OrderBy("-id").All(&result)

	if queryErr != nil {
		logs.Error(queryErr)
		return 0, result
	}
	count := []orm.Params{}
	media_os.ORM_CON.Raw("SELECT COUNT(title) as count FROM "+media_os.TABLE_NAME+where, valueSlice...).Values(&count)
	num, _ := strconv.Atoi(count[0]["count"].(string))
	return int64(num), result
}
func (this *MediaModel) ReadMediaOne(info MediaOptions) raw.Media {
	media_os._init()
	templateList := raw.Media{}
	if info.Id <= 0 {
		return templateList
	}

	error := media_os.OS.Filter("id", info.Id).One(&templateList)
	if error != nil {
		logs.Error(error)
		return templateList
	}

	return templateList
}

func (this *MediaModel) MediaModify(options raw.Media, fileds ...string) int64 {
	media_os._init()
	intRe, error := media_os.ORM_CON.Update(&options, fileds...)
	if error != nil {
		logs.Error(error)
		panic("error")
	}
	return intRe
}

func (this *MediaModel) MediaUpdateOrCreateById(id int, params raw.Media) (int64, error) {
	media_os._init()
	if id <= 0 {
		intRe, error := media_os.ORM_CON.Insert(&params)
		if error != nil {
			return 0, error
		}
		return intRe, nil
	}
	intRe, error := media_os.OS.Filter("id", id).Update(orm.Params{
		"title":           params.Title,
		"alias":           params.Alias,
		"foreign_name":    params.Foreign_name,
		"dubbing":         params.Dubbing,
		"director_ids":    params.Director_ids,
		"directors":       params.Directors,
		"main_actor_ids":  params.Main_actor_ids,
		"main_actors":     params.Main_actors,
		"writer_ids":      params.Writer_ids,
		"writers":         params.Writers,
		"rating":          params.Rating,
		"publish_company": params.Publish_company,
		"publish_year":    params.Publish_year,
		"region":          params.Region,
		"duration":        params.Duration,
		"language":        params.Language,
		"release_time":    params.Release_time,
		"category_id":     params.Category_id,
		"category_name":   params.Category_name,
		"type":            params.I__type,
		//"series_title":       params.Series_title,
		"desc":               params.Desc,
		"global_box_office":  params.Global_box_office,
		"chinese_box_office": params.Chinese_box_office,
		"brief_comment":      params.Brief_comment,
		"douban_score":       params.Douban_score,
		"maoyan_score":       params.Maoyan_score,
		"douban_url":         params.Douban_url,
		"status":             params.Status,
		"update_time":        time.Now().Unix(),
	})
	if error != nil {
		return 0, error
	}
	return intRe, nil
}

func (this *MediaModel) MediaInsert(options raw.Media) (int64, bool) {
	media_os._init()
	intRe, error := media_os.ORM_CON.Insert(&options)
	if error != nil {
		logs.Error(error)
		return 0, false
	}
	return intRe, true
}

func (this *MediaModel) ReadMediaById(ids []int) (int64, []raw.Media) {
	media_os._init()
	templateList := []raw.Media{}
	intRe, error := media_os.OS.Filter("id__in", ids).All(&templateList)
	if error != nil {
		logs.Error(error)
		return 0, templateList
	}
	return intRe, templateList
}
