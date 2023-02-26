package controllers

import (
	"example/demo1/initializer"
	"example/demo1/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	//Get data off req body

	var body struct {
		Title string
		Body  string
	}

	//here we bind our req data to body
	c.Bind(&body)

	//Create a post
	post := models.POST{Title: body.Title, Body: body.Body}

	result := initializer.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return it

	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostIndex(c *gin.Context) {
	//Get the posts

	// Get all records
	var post []models.POST
	initializer.DB.Find(&post)

	//respon d to them
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostShow(c *gin.Context) {
	//Get id from URL
	id := c.Param("id")

	//Get the post

	var post models.POST
	initializer.DB.First(&post, id)
	// SELECT * FROM users WHERE id = 10;

	//respon d to them
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostUpdate(c *gin.Context) {
	//Get the id of URL
	id := c.Param("id")

	//Get the data off req body from context
	var body struct {
		Title string
		Body  string
	}

	//here we bind our req data to body
	c.Bind(&body)

	//Find the post by id
	var post models.POST
	initializer.DB.First(&post, id)
	//updating

	initializer.DB.Model(&post).Updates(models.POST{Title: body.Title, Body: body.Body})

	//responed back
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsDelete(c *gin.Context) {
	// GEt the id of url
	id := c.Param("id")
	//delete the particular post
	initializer.DB.Delete(&models.POST{}, id)

	//respond back
	c.Status(200)
}
