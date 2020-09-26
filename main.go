package main

import (
	_ "BeegoProject001/db_mysql"
	_ "BeegoProject001/routers"
	"github.com/astaxie/beego"
)

func main() {
	//config:=beego.AppConfig
	//appname:=config.String("appname")
	//fmt.Println("应用名称",appname)
	//post,err:=config.Int("httppost")
	//if err!=nil {
	//	panic("编码成功")
	//
	//
	//}
	//fmt.Println(post)



	beego.Run()
}

