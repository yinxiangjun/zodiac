package routers

import (
	"xstv/cloud/cms-admin/admin/api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/*", &controllers.BaseController{}, "options:Options")
	beego.Router("/cms/category/lists", &controllers.CategoryController{}, "get:Lists")
	beego.Router("/cms/category/modify", &controllers.CategoryController{}, "post:Modify")
	beego.Router("/cms/category/delete", &controllers.CategoryController{}, "post:Delete")
	beego.Router("/cms/category/filter/list", &controllers.CategoryController{}, "get:FilterTypeList")
	beego.Router("/cms/category/filter/modify", &controllers.CategoryController{}, "post:FilterTypeModify")
	beego.Router("/cms/category/filter/delete", &controllers.CategoryController{}, "post:FilterTypeDelete")
	beego.Router("/cms/category/filteritem/list", &controllers.CategoryController{}, "get:FilterItemLists")
	beego.Router("/cms/topic/create", &controllers.TopicController{}, "post:Create;options:Options")
	beego.Router("/cms/topic/edit", &controllers.TopicController{}, "post:Edit;options:Options")
	beego.Router("/cms/filter/list", &controllers.FilterController{}, "get:List")
	beego.Router("/cms/topic/list", &controllers.TopicController{}, "get:TopicList")
	beego.Router("/cms/topic/additem", &controllers.TopicController{}, "post:Additem;options:Options")
	beego.Router("/cms/topic/videoreorder", &controllers.TopicController{}, "post:Videoreorder;options:Options")
	beego.Router("/cms/topic/removeitem", &controllers.TopicController{}, "post:Removeitem;options:Options")
	beego.Router("/cms/topic/detail", &controllers.TopicController{}, "get:Detail")
	beego.Router("/cms/topic/setstatus", &controllers.TopicController{}, "post:SetStatus;options:Options")
	beego.Router("/cms/topic/delete", &controllers.TopicController{}, "post:Delete;options:Options")
	beego.Router("/cms/category/filteritem/modify", &controllers.CategoryController{}, "post:FilterItemModify")
	beego.Router("/cms/category/filteritem/delete", &controllers.CategoryController{}, "post:FilterItemDelete")
	beego.Router("/cms/theme/list", &controllers.ThemeController{}, "get:List")
	beego.Router("/cms/theme/modify", &controllers.ThemeController{}, "post:Modify")
	beego.Router("/cms/theme/delete", &controllers.ThemeController{}, "post:Delete")
	beego.Router("/cms/theme/layout/save", &controllers.ThemeController{}, "post:LayoutSave")
	beego.Router("/cms/wrap/list", &controllers.WrapController{}, "get:List")
	beego.Router("/cms/wrap/modify", &controllers.WrapController{}, "post:Modify")
	beego.Router("/cms/wrap/add", &controllers.WrapController{}, "post:Add")
	beego.Router("/cms/wrap/remove", &controllers.WrapController{}, "post:Remove")
	beego.Router("/cms/wrap/layout/list", &controllers.WrapController{}, "get:LayoutList")
	beego.Router("/cms/wrap/layout/setup", &controllers.WrapController{}, "post:LayoutSetup")
	beego.Router("/cms/wrap/layout/add", &controllers.WrapController{}, "post:LayoutAdd")
	beego.Router("/cms/wrap/layout/remove", &controllers.WrapController{}, "post:LayoutRemove")
	beego.Router("/cms/wrap/layout/save", &controllers.WrapController{}, "post:LayoutSave")

	beego.Router("/cms/wrap/topic/list", &controllers.WrapController{}, "get:WrapTopicList")
	beego.Router("/cms/wrap/topic/add", &controllers.WrapController{}, "post:WrapTopicAdd")
	beego.Router("/cms/wrap/topic/modify", &controllers.WrapController{}, "post:WrapTopicModify")
	beego.Router("/cms/wrap/topic/remove", &controllers.WrapController{}, "post:WrapTopicRemove")

	beego.Router("/cms/theme/recommend/info", &controllers.ThemeController{}, "get:RecommendInfo")
	beego.Router("/cms/theme/recommend/modify", &controllers.ThemeController{}, "post:RecommendModify")
	beego.Router("/cms/theme/recommend/remove", &controllers.ThemeController{}, "post:RecommendRemove")
	beego.Router("/cms/theme/desktop/getthemeinfo", &controllers.ThemeController{}, "get:GetThemeInfo")
	beego.Router("/cms/theme/topic/list", &controllers.ThemeController{}, "get:ThemeTopicList")
	beego.Router("/cms/theme/topic/add", &controllers.ThemeController{}, "post:ThemeTopicAdd")
	beego.Router("/cms/theme/topic/modify", &controllers.ThemeController{}, "post:ThemeTopicModify")
	beego.Router("/cms/theme/topic/remove", &controllers.ThemeController{}, "post:ThemeTopicRemove")
	beego.Router("/cms/media/search", &controllers.MediaController{}, "get:List")
	beego.Router("/cms/recommend/lists",&controllers.RecommendController{},"get:Lists")
	beego.Router("/cms/recommend/detail", &controllers.RecommendController{}, "get:Detail")
	beego.Router("/cms/recommend/add", &controllers.RecommendController{}, "post:Add")
	beego.Router("/cms/recommend/remove", &controllers.RecommendController{}, "post:Remove")
	beego.Router("/cms/recommend/videoreorder", &controllers.RecommendController{}, "post:VideoReorder")

}
