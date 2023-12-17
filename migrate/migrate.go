package main

import (
	"github.com/jrab66/go-crud/initializers"
	"github.com/jrab66/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
