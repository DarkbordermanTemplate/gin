package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "practice/docs"
)

func Create_router() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Handle("GET", "/fruit", Get_by_query)
	router.Handle("GET", "/fruit/:name", Get_by_param)
	router.Handle("POST", "/fruit", Post_by_JSON)
	router.Handle("GET", "/health", Health)

	return router
}
