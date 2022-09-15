package middlewares

import (
	"github.com/gin-gonic/gin"
)

/**//**
 * @Author wangdt
 * @Description //TODO ${end}
 * @Date 2022-09-15 19:34:31
 * @Param x
 **/
func InitMiddleware(x *gin.Context) {
	x.Set("username", "zhangsan")
}
