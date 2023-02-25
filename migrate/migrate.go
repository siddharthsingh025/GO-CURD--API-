package main

import (
	"example/demo1/initializer"
	"example/demo1/models"
)

func init() {
	initializer.LoadEnvVariable()
	initializer.ConnectToDB()
}

func main() {
	// Migrate the schema
	initializer.DB.AutoMigrate(&models.POST{})

}
