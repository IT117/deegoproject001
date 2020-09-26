package controllers

import (
	"BeegoProject001/db_mysql"
	"BeegoProject001/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type RegidterController struct {
	beego.Controller
}

func (r *RegidterController) Post(){
	fmt.Println(r==nil)
	fmt.Println(r.Ctx ==nil)
	fmt.Println(r.Ctx.Request==nil)
	bodyBytes,err:=ioutil.ReadAll(r.Ctx.Request.Body)
	if err!=nil {
		r.Ctx.WriteString("数据接受错误，请重试")
		return

	}
	var user models.Usre
	err =json.Unmarshal(bodyBytes,&user)
	if err!=nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("数据解析错误")
		//3.将解析到的用户数据，保存到数据
		id ,err:=db_mysql.InserUser(user)
		if err!=nil {
			fmt.Println(err.Error())
			r.Ctx.WriteString("用户保存信息失败")
			return

		}
		fmt.Println(id)
		result :=models.ResponseResult{
			Code:    0,
			Message: "保存成功",
			Data:    nil,
		}
		r.Data["json"]=&result
		r.ServeJSON()

	}


}