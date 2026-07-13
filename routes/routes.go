package routes

import (
	"github.com/awanda/backend-repo/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(postHandler *handler.PostHandler) *gin.Engine {
	r := gin.Default()

	article := r.Group("/article")
	{
		article.POST("", postHandler.Create)
		article.GET("/:a/:b", postHandler.GetList)
		article.GET("/:a", postHandler.GetDetail)
		article.PUT("/:a", postHandler.Update)
		article.DELETE("/:a", postHandler.Delete)
	}

	return r
}
