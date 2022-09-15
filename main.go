package main

import (
	"ginproject/routers"
	"github.com/gin-gonic/gin"
)
import _ "net/http"

/**
 * @Author wangdt
 * @Date 2022-09-15 19:35:20
 * @Description: 
 */
type stuck struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

// 全局中间件
func initmiddlewaretwo(context *gin.Context) {

}

func initmiddlewareone(context *gin.Context) {

}

func main() {
	r := gin.Default()

	//加载路由文件-
	routers.ApiRoutersInit(r)

	//全局中间件
	r.Use(initmiddlewareone, initmiddlewaretwo)

	//启动
	r.Run()
}
