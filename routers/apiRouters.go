package routers

import (
	api "ginproject/controllers/api"
	"ginproject/middlewares"
	"github.com/gin-gonic/gin"
)

/**
 * @Author wangdt
 * @Description //TODO ${end}
 * @Date 2022-09-15 19:35:08
 * @Param r
 **/
func ApiRoutersInit(r *gin.Engine) {

	apiRouter := r.Group("/api",
		middlewares.InitMiddleware)
	{
		apiRouter.GET("user/index",
			api.UserController{}.Index)
		apiRouter.GET(
			"user/list",
			api.UserController{}.List)
		apiRouter.GET(
			"user/info",
			api.UserController{}.Info,
		)
		apiRouter.GET("user/get",
			api.UserController{}.Get)
		apiRouter.GET("user/set",
			api.UserController{}.Set)
	}
}
