package main

import (
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	"github.com/syahjamal/gin-full-api/config"
	"github.com/syahjamal/gin-full-api/middleware"
	"github.com/syahjamal/gin-full-api/routes"
)

func main() {
	//Setup Database
	config.InitDB()
	defer config.DB.Close()
	gotenv.Load()

	//Set up routing/router
	router := gin.Default()

	//Grouping router agar rapih dan jika ada perubahan mudah untuk trace
	v1 := router.Group("/api/v1/")
	{

		v1.GET("/auth/:provider", routes.RedirectHandler)
		v1.GET("/auth/:provider/callback", routes.CallbackHandler)

		//Halaman profil
		v1.GET("/profile", middleware.IsAuth(), routes.GetProfile)

		v1.GET("article/:slug", routes.GetArticle)
		articles := v1.Group("/articles")
		{
			articles.GET("/", routes.GetHome)
			articles.POST("/", middleware.IsAuth(), routes.PostArticle)
			articles.GET("tag/:tag", routes.GetArticleByTag)
			articles.PUT("/update/:id", middleware.IsAuth(), routes.UpdateArticle)

		}
	}
	router.Run()
}
