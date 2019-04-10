package utils

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"os"
)

var DbConfig config.Configer

type Response struct {
	Errno  int         `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data,omitempty"`
}

func LoadConf(dbConfigPath string) {
	var err error
	DbConfig, err = config.NewConfig("ini", dbConfigPath)
	if err != nil {
		fmt.Printf("read db config failed: %s\n", err)
		os.Exit(1)
	}
}

func GetDBConn() string {
	dbhost := DbConfig.DefaultString("host", "172.16.0.104:3306")
	dbuser := DbConfig.DefaultString("user", "root")
	dbpassword := DbConfig.DefaultString("password", "")
	db := DbConfig.DefaultString("db", "")
	conn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ")/" + db + "?charset=utf8&loc=Asia%2FShanghai"

	return conn
}
