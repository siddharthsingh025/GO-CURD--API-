package controllers

import (
	"example/demo1/initializer"
	"example/demo1/models"

	"github.com/gin-gonic/gin"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("gin-server")

func PostsCreate(c *gin.Context) {

	//span
	_, span := tracer.Start(c.Request.Context(), "PostsCreate")
	defer span.End()

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

	_, span := tracer.Start(c.Request.Context(), "PostIndex")
	defer span.End()

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

	_, span := tracer.Start(c.Request.Context(), "PostShow", oteltrace.WithAttributes(attribute.String("id", id)))
	defer span.End()

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

	_, span := tracer.Start(c.Request.Context(), "PostUpdate", oteltrace.WithAttributes(attribute.String("id", id)))
	defer span.End()

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

	_, span := tracer.Start(c.Request.Context(), "PostsDelete", oteltrace.WithAttributes(attribute.String("id", id)))
	defer span.End()
	//delete the particular post
	initializer.DB.Delete(&models.POST{}, id)

	//respond back
	c.Status(200)
}
