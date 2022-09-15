package middlewares

import (
	"github.com/gin-gonic/gin"
)

func InitMiddleware(x *gin.Context) {
	x.Set("username", "zhangsan")
}
