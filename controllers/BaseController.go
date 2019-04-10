package controllers

import (
	"cloud/zodiac/utils"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/goinggo/mapstructure"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	ZERO = iota
	THEME_TYPE
	WRAP_TYPE
)

type BaseController struct {
	beego.Controller
	response utils.Response
}

type _Init interface {
	_Init()
}

var UserInfo Login
var lsessionResult LsessionRes

type LsessionRes struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Data map[string]interface{}
}
type Login struct {
	UID         string `json:"uid" bson:"uid"`
	Username    string `json:"username" bson:"username"`
	Email       string `json:"email" bson:"email"`
	Realname    string `json:"realname" bson:"realname"`
	Logtime     int    `json:"logtime" bson:"logtime"`
	Permissions []struct {
		Act_name string `json:"act_name" bson:"act_name"`
		ID       string `json:"id" bson:"id"`
		Path     string `json:"path" bson:"path"`
	} `json:"permissions" bson:"permissions"`
}

func (this *BaseController) DealHeaders() {
	allowHeaders := "Qsc-Token,Content-Type,x-access-token,x-url-path"
	allowOrigin := "*"

	if this.Ctx.Input.Header("HTTP_ACCESS_CONTROL_REQUEST_HEADERS") != "" {
		allowHeaders = this.Ctx.Input.Header("HTTP_ACCESS_CONTROL_REQUEST_HEADERS")
	}

	if this.Ctx.Input.Header("Origin") != "" {
		allowOrigin = this.Ctx.Input.Header("Origin")
	}

	this.Ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
	this.Ctx.Output.Header("Access-Control-Allow-Origin", allowOrigin)
	this.Ctx.Output.Header("Access-Control-Allow-headers", allowHeaders)
	this.Ctx.Output.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
	this.Ctx.Output.Header("Access-Control-Request-Method", "*")
	this.Ctx.Output.Header("Access-Control-Request-Headers", "*")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	this.Ctx.Output.Header("Access-Control-Expose-Headers", "X-My-Custom-Header, X-Another-Custom-Header")
}

func (this *BaseController) Options() {
	this.DealHeaders()
	this.Ctx.Output.Header("HTTP/1.1 200 No Content", "")
	return
}

func (this *BaseController) ServeJSON(encoding ...bool) {
	this.DealHeaders()
	this.Controller.ServeJSON(encoding...)
}

func (this *BaseController) Catch() {
	if r := recover(); r != nil {
		this.response.Errno = CODE_PARSE_DATA_FAIL
		this.response.Errmsg = fmt.Sprintf("%s", r)
		this.LogDebuger(r)
		this.Data["json"] = this.response
		this.ServeJSON()
	}
}

func (this *BaseController) FormatTrace(r interface{}) string {
	return fmt.Sprintf(
		"[%s] [%s] [%s] [%s]",
		UserInfo.Email,
		this.Ctx.Request.RequestURI,
		this.Ctx.Input.RequestBody,
		fmt.Sprintf("%s", r),
	)
}

func (this *BaseController) LogDebuger(r interface{}) {
	logs.Error(this.FormatTrace(r) + " [--MESSAGE--]: " + strings.Replace(string(debug.Stack()), "\n", "", -1))
}

func (this *BaseController) _init() {
	this.response = utils.Response{
		Errno:  CODE_OK,
		Errmsg: MSG_OK,
	}
	this.Data["json"] = this.response
}

var lessionId string

func (this *BaseController) Prepare() {
	if app, ok := this.AppController.(_Init); ok {
		app._Init()
	}
	// if this.Ctx.Request.Method == http.MethodOptions {
	// 	return
	// }
	// nolog := utils.DbConfig.String("AUTHORIZATION_NOLOGIN")
	// if nolog == "true" {
	// 	UserInfo.Email = ""
	// 	UserInfo.Username = ""
	// 	UserInfo.UID = ""
	// 	UserInfo.Logtime = 0
	// 	UserInfo.Realname = ""
	// 	return
	// }
	// lessionId = this.Ctx.GetCookie("lsession")
	// if !bson.IsObjectIdHex(lessionId) {
	// 	this.reLogin()
	// }
	// this.getUserinfo(lessionId)
}

func (this *BaseController) getUserinfo(lessionId string) {
	var mongoSession *mgo.Session
	mongoSession, _ = mgo.Dial(utils.DbConfig.String("mongo::uri"))
	conn := mongoSession.DB(utils.DbConfig.String("mongo::db")).C(utils.DbConfig.String("mongo::collection"))
	err := conn.Find(bson.M{"_id": bson.ObjectIdHex(lessionId)}).One(&lsessionResult)
	if err != nil {
		logs.Error(err.Error())
		this.reLogin()
	}
	err = mapstructure.Decode(lsessionResult.Data["login"], &UserInfo)
	if err != nil {
		logs.Error(err.Error())
	}
	if UserInfo.Email == "" || UserInfo.Username == "" {
		logs.Error("lession is illegal, Username=" + UserInfo.Username + ", Email=" + UserInfo.Email)
		this.reLogin()
	}
	logs.Info("[" + UserInfo.Email + "] [--access the system--] (" + this.FormatTrace(UserInfo.UID) + ")")
}

func (this *BaseController) SetLsession(key string, value interface{}) {
	var mongoSession *mgo.Session
	mongoSession, _ = mgo.Dial(utils.DbConfig.String("mongo::uri"))
	conn := mongoSession.DB(utils.DbConfig.String("mongo::db")).C(utils.DbConfig.String("mongo::collection"))
	lsessionResult.Data[key] = value
	conn.Update(bson.M{"_id": bson.ObjectIdHex(lessionId)}, bson.M{"$set": bson.M{"data." + key: value}})
	logs.Info(fmt.Sprintf("set lsession: key=%s value=%v", key, value))
}

func (this *BaseController) GetLsession(key string) interface{} {
	data := lsessionResult.Data[key]
	return data
}

func (this *BaseController) reLogin() {
	http.Redirect(this.Ctx.ResponseWriter, this.Ctx.Request, utils.DbConfig.String("domain::web")+"/login", 302)
	return
}

type bodyStruct struct {
	ID                 int64    `json:"id"`
	Title              string   `json:"title"`
	Foreign_name       string   `json:"foreignName"`
	ForeignName        string   `json:"foreign_name"`
	Rating             string   `json:"rating"`
	Release_company    string   `json:"company"`
	Release_year       int      `json:"year"`
	Region             string   `json:"Region"`
	Duration           int      `json:"duration"`
	Language           string   `json:"language"`
	Official_website   string   `json:"website"`
	Release_time       int      `json:"releaseTime"`
	Series_name        string   `json:"seriesName"`
	Global_box_office  int      `json:"globalBoxOffice"`
	Chinese_box_office int      `json:"chineseBoxOffice"`
	Story_source       int      `json:"storySource"`
	Is_remade          int      `json:"isRemade"`
	Brief_comment      string   `json:"comment"`
	Desc               string   `json:"desc"`
	Type               string   `json:"type"`
	Weight             int      `json:"weight"`
	Play_count         int      `json:"playCount"`
	Origin_score       float64  `json:"originScore"`
	Douban_score       float64  `json:"doubanScore"`
	Maoyan_score       float64  `json:"maoyanScore"`
	Imdb_score         float64  `json:"imdbScore"`
	Douban_url         string   `json:"doubanUrl"`
	Maoyan_url         string   `json:"maoyanUrl"`
	Imdb_url           string   `json:"imdbUrl"`
	Heat               int      `json:"heat"`
	Praises            int      `json:"praises"`
	Notes              string   `json:"notes"`
	Name               string   `json:"name"`
	Status             int      `json:"status"`
	Alias              string   `json:"alias"`
	Code               string   `json:"code"`
	Level              int      `json:"level"`
	Start_time         int64    `json:"startTime"`
	End_time           int64    `json:"endTime"`
	Width              int      `json:"width"`
	Height             int      `json:"height"`
	Ratio              int      `json:"ratio"`
	AlbumId            int64    `json:"albumId"`
	IsModify           int      `json:"isModify"`
	Person_ids         []int    `json:"personIds"`
	Person_id          int      `json:"personId"`
	PersonId           int      `json:"person_id"`
	Tag_ids            []int    `json:"tagIds"`
	Tag_id             int      `json:"tagId"`
	Person_type        int      `json:"personType"`
	Rank               int      `json:"rank"`
	Logo               string   `json:"logo"`
	Best_resolution    string   `json:"resolution"`
	Category_id        int      `json:"categoryId"`
	Category_name      string   `json:"categoryName"`
	Play_feature       string   `json:"feature"`
	Xstv_status        int      `json:"xstvStatus"`
	XstvStatus         int      `json:"xstv_status"`
	Src_code           string   `json:"src_code"`
	Episode_id         int      `json:"episode_id"`
	Content_type       int      `json:"content_type"`
	PosterUrls         []string `json:"posterUrls"`
	Portrait           string   `json:"portrait"`
	// ForeignName        string   `json:"foreign_name"`
	Gender        int    `json:"gender"`
	Birthday      int    `json:"birthday"`
	Deathday      int    `json:"deathday"`
	Birthplace    string `json:"birthplace"`
	Citizenship   string `json:"citizenship"`
	Occupation    string `json:"occupation"`
	Nationality   string `json:"nationality"`
	Blood_group   string `json:"blood_group"`
	Constellation string `json:"constellation"`
	Relation      string `json:"relation"`
	Debut_time    int    `json:"debut_time"`
	Hobby         string `json:"hobby"`
}
