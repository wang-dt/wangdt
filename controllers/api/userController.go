package api

import (
	"fmt"
	"ginproject/models"
	"github.com/gin-gonic/gin"
)

/**
 * @Author wangdt
 * @Date 2022-09-15 19:34:15
 * @Description:
 */
type UserController struct {
}

func (c UserController) Info(r *gin.Context) {

	//从中间件中获取数据
	var username, _ = r.Get("username")
	r.JSON(200, gin.H{
		"username": username,
	})

	//使用公共方法
	ms := models.UnixTotime(1663150354)
	r.String(200, ms)
	r.String(200, "UserController")

}

func (c UserController) Index(r *gin.Context) {

}

// 列表- 查询所有用户
func (c UserController) List(r *gin.Context) {
	userList := []models.User{}
	models.DB.Find(&userList)
	r.JSON(200, gin.H{
		"result": userList,
	})
}

// 获取一条数据
func (c UserController) Get(r *gin.Context) {
	fmt.Println("----userinfo")
}

// 添加数据
func (c UserController) Set(r *gin.Context) {
	m := models.User{
		Username: "王武",
		Age:      "19",
		Address:  "哈了三舅",
	}
	models.DB.Create(&m)
}
