package db_mysql

import (
	"BeegoProject001/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/go-sql-driver/mysql"

)

var DB *sql.DB

func init() {
	fmt.Println("链接数据库")
	config := beego.AppConfig
	dbDriver := config.String("db_driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	connUrl := dbUser + ":" + dbPassword + "@tcp(" + dbIp + ")/" + dbName + "?charset=utf8"
	db, err := sql.Open(dbDriver, connUrl)
	if err != nil {
		fmt.Println(err.Error())
	}
	DB = db


}
func InserUser(usre models.Usre) (int64,error) {
	hashMd5:=md5.New()
hashMd5.Write([]byte(usre.Password))
	bytes:=hashMd5.Sum(nil)
	usre.Password =hex.EncodeToString(bytes)
	fmt.Println("将要保存的用户名：",usre.Nick,"密码",usre.Password)
	result,err :=DB.Exec("insert into beego_test (nick,passwork) value (?,?)",usre.Nick,usre.Password)
	if err!=nil{
		return -1, err
	}
	id,err:=result.RowsAffected()
	if err!=nil {
		return -1,err

	}
	return id,nil
}
func QueryUser(){
	DB.QueryRow("select *from")
}