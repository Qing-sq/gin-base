package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"time"
)
var (
	DB *gorm.DB
)

func InitDB()  {
	dsn:= viper.GetString("mysql.dsn")
	db,err := gorm.Open("mysql",dsn)
	if err != nil{
		fmt.Println(err,"数据库连接失败")
	}
	//默认不加复数
	db.SingularTable(true)
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(viper.GetInt("max_idle"))
	//打开
	db.DB().SetMaxOpenConns(viper.GetInt("max_open"))
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	DB.AutoMigrate(&User{}).AutoMigrate(&Good{})
}
