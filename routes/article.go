package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/syahjamal/gin-full-api/config"
	"github.com/syahjamal/gin-full-api/models"
)

func GetHome(c *gin.Context) {

	items := []models.Article{}
	config.DB.Find(&items)

	c.JSON(200, gin.H{
		"status": "berhasil ke halaman home",
		"data":   items,
	})
}

func GetArticle(c *gin.Context) {
	//parameter
	slug := c.Param("slug")

	var item models.Article

	//Query di gorm = Select * from table where slug = "slug"
	if config.DB.First(&item, "slug = ?", slug).RecordNotFound() {
		c.JSON(404, gin.H{"status": "error", "message": "record not found"})
		c.Abort() //Batalin request
		return
	}

	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   item,
	})
}

func PostArticle(c *gin.Context) {
	item := models.Article{
		Title: c.PostForm("title"),
		Desc:  c.PostForm("desc"),
		Slug:  slug.Make(c.PostForm("title")),
	}

	//Mencegah slug sama, maka generate random slug

	config.DB.Create(&item)

	c.JSON(200, gin.H{
		"status": "berhasil post",
		"data":   item,
	})
}
