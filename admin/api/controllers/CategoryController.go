package controllers

import (
	"bytes"
	"encoding/json"
	"strings"
	"xstv/cloud/cms-admin/admin/api/models"
	"xstv/cloud/cms-admin/common/raw"
)

type CategoryController struct {
	BaseController
	models.CategoryModel
}

type CategoryRows struct {
	List     []raw.Category `json:"list"`
	Total    int64          `json:"total"`
	Page     int            `json:"page"`
	PageSize int            `json:"pageSize"`
}

type FilterTypeRow struct {
	List     []raw.CategoryFiltertype `json:"list"`
	Total    int64                    `json:"total"`
	Page     int                      `json:"page"`
	PageSize int                      `json:"page_size"`
}

type FilterItemRow struct {
	List     []CategoryFilteritemList `json:"list"`
	Total    int64                    `json:"total"`
	Page     int                      `json:"page"`
	PageSize int                      `json:"page_size"`
}

type CategoryFilteritemList struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Sort      int    `json:"sort"`
	Desc      string `json:"desc"`
	Condition []Condions `json:"condition"`
	Filterid  int64  `json:"filterid"`
}


func (self *CategoryController) Lists() {
	self._init()
	defer self.Catch()

	name := self.GetString("name")
	Type, _ := self.GetInt("type")
	page, _ := self.GetInt("page", 1)
	pagesize, _ := self.GetInt("pagesize", 30)
	oparams := models.CategoryOptions{
		Name:   name,
		Type:   Type,
		Limit:  pagesize,
		Offset: (page - 1) * pagesize,
	}
	count, Lists := self.CategoryInfo(oparams)
	List := CategoryRows{
		Total:    count,
		Page:     page,
		PageSize: pagesize,
	}
	if len(Lists) == 0 {
		List.List = []raw.Category{}
		self.response.Data = List
		self.Data["json"] = self.response
		self.ServeJSON()
	}
	List.List = Lists
	self.response.Data = List
	self.Data["json"] = self.response
	self.ServeJSON()
}

func (self *CategoryController) Modify() {
	self._init()
	defer self.Catch()

	bodyInfo := raw.Category{}
	requestJson := self.Ctx.Input.RequestBody
	err := json.Unmarshal(requestJson, &bodyInfo)
	if err != nil {
		panic(err)
	}

	if bodyInfo.Status <= 0 {
		bodyInfo.Status = 1
	}

	bodyInfo.Name = strings.TrimSpace(bodyInfo.Name)
	if bodyInfo.Name == "" {
		panic("分类名称字段格式有误！")
	}
	if bodyInfo.Desc == "" {
		panic("描述字段格式有误！")
	}
	params := raw.Category{
		Name:   bodyInfo.Name,
		Desc:   bodyInfo.Desc,
		Type:   2,
		Status: 1,
	}

	mid := self.CategoryModify(params, bodyInfo.Id)
	self.response.Data = map[string]int64{"id": mid}
	self.Data["json"] = self.response
	self.ServeJSON()
}

func (self *CategoryController) Delete() {
	self._init()
	defer self.Catch()

	bodyInfo := raw.Category{}
	requestJson := self.Ctx.Input.RequestBody
	err := json.Unmarshal(requestJson, &bodyInfo)
	if err != nil {
		panic(err)
	}
	if bodyInfo.Id <= 0 {
		panic("分类ID格式有误！")
	}
	self.CategoryDelete(bodyInfo.Id)
	self.Data["json"] = self.response
	self.ServeJSON()
}

func (self *CategoryController) FilterTypeList() {
	self._init()
	defer self.Catch()
	categoryID, _ := self.GetInt64("id")
	if categoryID <= 0 {
		panic("分类ID格式有误！")
	}
	page, _ := self.GetInt("page", 1)
	pagesize, _ := self.GetInt("pagesize", 30)
	oparams := models.FilterType{
		Id:     categoryID,
		Limit:  pagesize,
		Offset: (page - 1) * pagesize,
	}
	count, Lists := self.FilterTypeInfo(oparams)
	List := FilterTypeRow{
		Total:    count,
		Page:     page,
		PageSize: pagesize,
	}
	if len(Lists) == 0 {
		List.List = []raw.CategoryFiltertype{}
		self.response.Data = List
		self.Data["json"] = self.response
		self.ServeJSON()
	}
	List.List = Lists
	self.response.Data = List
	self.Data["json"] = self.response
	self.ServeJSON()
}

func (self *CategoryController) FilterTypeModify() {
	self._init()
	defer self.Catch()

	bodyInfo := raw.CategoryFiltertype{}
	requestJson := self.Ctx.Input.RequestBody
	err := json.Unmarshal(requestJson, &bodyInfo)
	if err != nil {
		panic(err)
	}
	if bodyInfo.Categoryid <= 0 {
		panic("分类ID字段格式有误！")
	}

	bodyInfo.Name = strings.TrimSpace(bodyInfo.Name)
	if bodyInfo.Name == "" {
		panic("分类名称字段格式有误！")
	}
	if bodyInfo.Desc == "" {
		panic("描述字段格式有误！")
	}
	params := raw.CategoryFiltertype{
		Name:       bodyInfo.Name,
		Desc:       bodyInfo.Desc,
		Sort:       bodyInfo.Sort,
		Categoryid: bodyInfo.Categoryid,
	}

	mid := self.CategoryFilterTypeModify(params, bodyInfo.Id)
	self.response.Data = map[string]int64{"id": mid}
	self.Data["json"] = self.response
	self.ServeJSON()
}

func (self *CategoryController) FilterTypeDelete() {
	self._init()
	defer self.Catch()

	bodyInfo := raw.CategoryFiltertype{}
	requestJson := self.Ctx.Input.RequestBody
	err := json.Unmarshal(requestJson, &bodyInfo)
	if err != nil {
		panic(err)
	}
	if bodyInfo.Id <= 0 {
		panic("分类筛选类型ID格式有误！")
	}
	self.CategoryFilterTypeDelete(bodyInfo.Id)
	self.Data["json"] = self.response
	self.ServeJSON()
}

func (self *CategoryController) FilterItemLists() {
	//
	self._init()
	defer self.Catch()

	filterID, _ := self.GetInt64("id")
	if filterID <= 0 {
		panic("筛选类型ID格式有误！")
	}
	page, _ := self.GetInt("page", 1)
	pagesize, _ := self.GetInt("pagesize", 30)
	oparams := models.FilterType{
		Id:     filterID,
		Limit:  pagesize,
		Offset: (page - 1) * pagesize,
	}
	count, Lists := self.FilterItemInfo(oparams)
	List := FilterItemRow{
		Total:    count,
		Page:     page,
		PageSize: pagesize,
	}
	if len(Lists) == 0 {
		List.List = []CategoryFilteritemList{}
		self.response.Data = List
		self.Data["json"] = self.response
		self.ServeJSON()
	}
	CateFilterInfo := []CategoryFilteritemList{}
	var conditions []Condions
	for _,value:=range Lists {
		json.Unmarshal([]byte(value.Condition), &conditions)
		if len(conditions) == 0 {
			conditions = []Condions{}
		}
		CateFilterItem := CategoryFilteritemList{
			Id:   value.Id,
			Name:         value.Name,
			Sort:        value.Sort,
			Desc:        value.Desc,
			Filterid:       value.Filterid,
			Condition:     conditions,
		}
		CateFilterInfo = append(CateFilterInfo, CateFilterItem)
	}
	List.List = CateFilterInfo
	self.response.Data = List
	self.Data["json"] = self.response
	self.ServeJSON()
}

func (self *CategoryController) FilterItemModify() {
	self._init()
	defer self.Catch()
	bodyInfo := CategoryFilteritemList{}
	requestJson := self.Ctx.Input.RequestBody
	err := json.Unmarshal(requestJson, &bodyInfo)
	if err != nil {
		panic(err)
	}
	if bodyInfo.Filterid <= 0 {
		panic("分类筛选类型ID格式有误！")
	}

	bodyInfo.Name = strings.TrimSpace(bodyInfo.Name)
	if bodyInfo.Name == "" {
		panic("分类筛选项名称格式有误！")
	}
	if bodyInfo.Desc == "" {
		panic("分类筛选项描述字段格式有误！")
	}

	conditions := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(conditions)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.Encode(bodyInfo.Condition)

	params := raw.CategoryFilteritem{
		Name:      bodyInfo.Name,
		Desc:      bodyInfo.Desc,
		Sort:      bodyInfo.Sort,
		Filterid:  bodyInfo.Filterid,
		Condition: conditions.String(),
	}

	mid := self.CategoryFilterItemModify(params, bodyInfo.Id)
	self.response.Data = map[string]int64{"id": mid}
	self.Data["json"] = self.response
	self.ServeJSON()
}

func (self *CategoryController) FilterItemDelete() {
	self._init()
	defer self.Catch()

	bodyInfo := raw.CategoryFilteritem{}
	requestJson := self.Ctx.Input.RequestBody
	err := json.Unmarshal(requestJson, &bodyInfo)
	if err != nil {
		panic(err)
	}
	if bodyInfo.Id <= 0 {
		panic("分类筛选项ID格式有误！")
	}
	self.CategoryFilterItemDelete(bodyInfo.Id)
	self.Data["json"] = self.response
	self.ServeJSON()
}
