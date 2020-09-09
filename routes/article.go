package routes

import (
	"strconv"
	"time"

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

	var oldItem models.Article
	slug := slug.Make(c.PostForm("title"))

	if !config.DB.First(&oldItem, "slug = ?", slug).RecordNotFound() {
		//generate string baru
		slug = slug + "-" + strconv.FormatInt(time.Now().Unix(), 10)
	}

	item := models.Article{
		Title:  c.PostForm("title"),
		Desc:   c.PostForm("desc"),
		Tag:    c.PostForm("tag"),
		Slug:   slug,
		UserID: uint(c.MustGet("jwt_user_id").(float64)),
	}

	//Mencegah slug sama, maka generate random slug
	config.DB.Create(&item)

	c.JSON(200, gin.H{
		"status": "berhasil post",
		"data":   item,
	})
}

func GetArticleByTag(c *gin.Context) {
	tag := c.Param("tag")
	items := []models.Article{}

	config.DB.Where("tag LIKE ?", "%"+tag+"%").Find(&items)

	c.JSON(200, gin.H{"data": items})

}

func UpdateArticle(c *gin.Context) {
	//parameter
	id := c.Param("id")

	var item models.Article

	//Query di gorm = Select * from table where slug = "slug"
	if config.DB.First(&item, "id = ?", id).RecordNotFound() {
		c.JSON(404, gin.H{"status": "error", "message": "record not found"})
		c.Abort() //Batalin request
		return
	}

	config.DB.Model(&item).Where("id = ?", id).Updates(models.Article{
		Title: c.PostForm("title"),
		Desc:  c.PostForm("desc"),
		Tag:   c.PostForm("tag"),
	})

	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   item,
	})
}
