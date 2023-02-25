package controllers

import (
	"example/demo1/initializer"
	"example/demo1/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	//Get data off req body

	//Create a post
	post := models.POST{Title: "Jinzhu", Body: "post body"}

	result := initializer.DB.Create(&post)

	if result != nil {
		c.Status(400)
		return
	}

	//Return it

	c.JSON(200, gin.H{
		"posts": post,
	})
}
