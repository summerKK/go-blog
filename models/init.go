package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbHost := beego.AppConfig.String("db::host")
	dbPort := beego.AppConfig.String("db::port")
	dbUser := beego.AppConfig.String("db::user")
	dbPassword := beego.AppConfig.String("db::password")
	dbName := beego.AppConfig.String("db::name")
	timezone := beego.AppConfig.String("timezone")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&loc=" + timezone

	orm.RegisterDataBase("default", "mysql", dsn)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}
