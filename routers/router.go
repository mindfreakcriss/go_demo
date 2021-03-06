package routers

import (
	_ "gin-blog/docs" //不加会内部错误
	"gin-blog/middleware/jwt"
	"gin-blog/pkg/setting"
	"gin-blog/routers/api"
	v1 "gin-blog/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	
	r.Use(gin.Logger())
	
	r.Use(gin.Recovery())
	
	gin.SetMode(setting.RunMode)

	//不加会404
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/auth",api.GetAuth)

	apiv1 := r.Group("/api/v1")
	//中间件使用
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}