package controllers

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/astaxie/beego"
)

type ZodiacController struct {
	beego.Controller
}

func (c *ZodiacController) OneDayFortuneList() {

	runtime.GC()
	fmt.Println(runtime.NumCPU())
	var fl float64
	var err error
	fl, err = strconv.ParseFloat("3.5", 64)
	fmt.Println(fl, err)
	fmt.Println(strconv.Itoa(5))
	fmt.Println(strconv.FormatBool(false))
	fmt.Println(strconv.Quote(`bs'd\s"d`))

	//data := []byte("abcdelckfdj")
	// has := md5.Sum(data)
	// md5str1 := fmt.Sprintf("%x", has)
	//fmt.Println(md5str1)

	//fmt.Println(string(debug.Stack()))
	//debug.PrintStack()

	c.Data["Website"] = "beego.me"

	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "oneday/index.tpl"
}

// func (self *TopicController) EditMoen() {
// 	defer self.Catch()

// 	topicId, err := self.GetInt("topic_id")
// 	if err != nil {
// 		panic("topic_id格式有误！")
// 	}
// 	var body TopicCreateBody
// 	bodyContent := self.Ctx.Input.RequestBody
// 	err = json.Unmarshal(bodyContent, &body)
// 	if err != nil {
// 		self.Abort(err.Error())
// 	}
// 	o := orm.NewOrm()
// 	filters := []raw.Filters{}
// 	o.QueryTable("filters").All(&filters)

// 	conditions := bytes.NewBuffer([]byte{})
// 	jsonEncoder := json.NewEncoder(conditions)
// 	jsonEncoder.SetEscapeHTML(false)
// 	jsonEncoder.Encode(body.Conditions)

// 	topic := raw.Topic{
// 		Id:         topicId,
// 		Name:       body.Name,
// 		Title:      body.Title,
// 		Desc:       body.Desc,
// 		CategoryId: body.CategoryId,
// 		Order:      body.Order,
// 		Condition:  conditions.String(),
// 		Status:     body.Status,
// 		OriginSize: body.OriginSize,
// 		Super:      body.IsSuper,
// 		UpdateTime: time.Now().Unix(),
// 	}
// 	mid, err := o.Update(&topic)
// 	if err != nil {
// 		logs.Error(err)
// 	}
// 	self.response.Data = mid
// 	self.Data["json"] = self.response
// 	self.ServeJSON()
// }
