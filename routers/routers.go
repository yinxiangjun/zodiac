package routers

import (
	"cloud/zodiac/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.ZodiacController{})
	//beego.Router("/*", &controllers.BaseController{}, "options:Options")
	beego.Router("/weixin/pay", &controllers.WeixinPayController{}, "get:WxPay")
	beego.Router("/zodiac/list", &controllers.ZodiacController{}, "get:OneDayFortuneList")
	beego.Router("/*", &controllers.BaseController{}, "options:Options")
}
