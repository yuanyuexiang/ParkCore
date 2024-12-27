package main

//package park

import (
	_ "park/routers"

	beeLogger "github.com/beego/bee/v2/logger"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "github.com/go-sql-driver/mysql"
)

func mainX() {
	err := orm.RegisterDriver("sqlite3", orm.DRSqlite)
	if err != nil {
		beeLogger.Log.Fatalf("%s", err)
	}
	err = orm.RegisterDataBase("default", "sqlite3", "./database/park.db")
	if err != nil {
		beeLogger.Log.Fatalf("%s", err)
	}
	orm.SetMaxOpenConns("default", 30)
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		beeLogger.Log.Fatalf("%s", err)
	}
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/park/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.Run()
}

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
		orm.Debug = false
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/park/swagger"] = "swagger"
	}
	beego.Run()
}

var running bool = false

func Run(dir string) {
	if running {
		return
	}

	beego.LoadAppConfig("ini", dir+"/conf/app.conf")
	err := orm.RegisterDriver("sqlite3", orm.DRSqlite)
	if err != nil {
		beeLogger.Log.Fatalf("%s", err)
	}
	err = orm.RegisterDataBase("default", "sqlite3", dir+"/database/park.db")
	if err != nil {
		beeLogger.Log.Fatalf("%s", err)
	}
	orm.SetMaxOpenConns("default", 30)
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		beeLogger.Log.Fatalf("%s", err)
	}
	if beego.BConfig.RunMode == "dev" {
		//orm.Debug = true
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir[dir+"/park/swagger"] = "swagger"
	}
	running = true
	beego.Run()
}

func Stop() {

}
