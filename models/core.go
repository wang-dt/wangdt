package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

/**
 * @Author wangdt
 * @Description //TODO ${end}
 * @Date 2022-09-15 19:34:45
 **/
func init() {

	dsn := "root:@tcp(127.0.0.1:3306)/springboot?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
}
