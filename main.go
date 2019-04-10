package main

import (
	"cloud/common/log"
	_ "cloud/zodiac/routers"
	"cloud/zodiac/utils"
	"flag"
	"os"
	"runtime"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var confFile = flag.String("conf", "conf/db.beta.conf", "config file of mysql(host|user|pwd...)")

func init() {
	//beego.BConfig.WebConfig.Session.SessionOn = true
	logs.SetLogger(logs.AdapterFile, `{"filename":"/xstv/app/media-admin-api-v2.log","daily":true,"maxdays":10}`)
}

func main() {
	flag.Parse()
	utils.LoadConf(*confFile)

	log.InitLogger("")
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	err := orm.RegisterDataBase("default", "mysql", utils.GetDBConn())
	if err != nil {
		//log.GLog.Error("RegisterDataBase:%s failed,err:%s", utils.GetDBConn(), err.Error())
		time.Sleep(3 * time.Second)
		os.Exit(1)
	}
	if beego.AppConfig.String("debugorm") == "1" {
		orm.Debug = true
	}
	beego.Run()
}
