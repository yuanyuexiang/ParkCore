package main

//package park

import (
	_ "park/routers"

	beeLogger "github.com/beego/bee/v2/logger"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	sqlConn, err := beego.AppConfig.String("sqlconn")
	if err != nil {
		beeLogger.Log.Fatal(err.Error())
	}
	err = orm.RegisterDataBase("default", "mysql", sqlConn)
	if err != nil {
		beeLogger.Log.Fatalf("%s", err)
	}
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		beeLogger.Log.Fatalf("%s", err)
	}
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/park/swagger"] = "swagger"
	}
	beego.Run()
}
