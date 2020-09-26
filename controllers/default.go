package controllers

import (
	"BeegoProject001/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

//func (c *MainController) Get() {
//	uersName := c.Ctx.Input.Query("use")
//	passWord := c.Ctx.Input.Query("psd")
//	if uersName != "admin" || passWord != "123456" {
//
//		c.Ctx.WriteString("信息错误")
//
//	} else {
//		c.Ctx.WriteString("成功输入")
//	}
//	c.Data["Website"] = "www.baidu.com"
//	c.Data["Email"] = "2714257837@qq.com"
//	c.TplName = "index.tpl"
//}
//func (c *MainController) Post() {
//	//接受post请求数据
//	name := c.Ctx.Request.FormValue("name")
//	age := c.Ctx.Request.FormValue("age")
//	sex := c.Ctx.Request.FormValue("sex")
//	fmt.Println(name, age, sex)
//	if name != "wangrenzhi" && age != "18" && sex != "maie" {
//		c.Ctx.WriteString("数据验证失败")
//		return
//
//	}
//	c.Ctx.WriteString("验证数据成功")
//
//}
//func (c *MainController) Post() {
//	//解析前端提交的json格式的数据
//	var person models.Preson
//	dataBytes, err := ioutil.ReadAll(c.Ctx.Request.Body)
//	if err != nil {
//		c.Ctx.WriteString("数据接受错误，请重试")
//		return
//
//	}
//	err = json.Unmarshal(dataBytes, &person)
//	if err != nil {
//		c.Ctx.WriteString("数据解析失败，请重试")
//		return
//	}
//	fmt.Println("姓名", person.Name)
//	fmt.Println("年龄", person.Age)
//	fmt.Println("性别", person.Sex)
//	c.Ctx.WriteString("数据解析成功")
//
//}
func (c *MainController) Post() {
	var person1 models.Preson
	//请求前端body内容
	dataBytes1,err :=ioutil.ReadAll(c.Ctx.Request.Body)
	if err!=nil {
		c.Ctx.WriteString("数据接受错误，请重试")
		return

	}
	//解析前端返回的josn内容
	err = json.Unmarshal(dataBytes1,&person1)
	if err!=nil {
		c.Ctx.WriteString("数据解析失败，请重试")
		return

	}
	fmt.Println("姓名",person1.Name)
	fmt.Println("生日",person1.Birthday)
	fmt.Println("地址",person1.Address)
	fmt.Println("昵称",person1.Nick)
	c.Ctx.WriteString("数据解析成功")
}
func(c *MainController)Delete(){

}
