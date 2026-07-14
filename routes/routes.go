package routes

import (
	"time"

	"github.com/awanda/backend-repo/internal/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(postHandler *handler.PostHandler) *gin.Engine {
	r := gin.Default()

	// CORS Configuration
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
		},
		MaxAge: 12 * time.Hour,
	}))

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