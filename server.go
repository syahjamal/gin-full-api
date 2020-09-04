package main

import (
	"github.com/gin-gonic/gin"
	"github.com/syahjamal/gin-full-api/config"
	"github.com/syahjamal/gin-full-api/routes"
)

func main() {
	//Setup Database
	config.InitDB()
	defer config.DB.Close()

	//Set up routing/router
	router := gin.Default()

	//Grouping router agar rapih dan jika ada perubahan mudah untuk trace
	v1 := router.Group("/api/v1/")
	{
		article := v1.Group("/article")
		{
			article.GET("/", routes.GetHome)
			article.GET("/:slug", routes.GetArticle)
			article.POST("/", routes.PostArticle)
		}
	}
	router.Run()
}
