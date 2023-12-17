// package main

// import "fmt"

//	func main() {
//		fmt.Printf("hello")
//	}
package main

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/jrab66/go-poc/controllers"
	"github.com/jrab66/go-poc/initializers"
	"github.com/jrab66/go-poc/models"
)

func getClientIP(c *gin.Context) string {
	// Check X-Real-IP or X-Forwarded-For headers for proxy scenarios
	if ip := c.Request.Header.Get("X-Real-IP"); ip != "" {
		return ip
	} else if forwarded := c.Request.Header.Get("X-Forwarded-For"); forwarded != "" {
		ips := strings.Split(forwarded, ",")
		return strings.TrimSpace(ips[0])
	}

	// Fallback to the default ClientIP method
	return c.ClientIP()
}

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	// migrator := initializers.DB.Migrator()
	// migrator.DropTable(&models.Post{})
	initializers.DB.AutoMigrate(&models.Post{})

}

func setup() {
	// initializers.LoadEnvVariables()
	// initializers.ConnectToDB()
	migrator := initializers.DB.Migrator()
	migrator.DropTable(&models.Post{})
	// database.Connect()
	// initializers.DB.AutoMigrate(&models.Post{})

	// database.Database.AutoMigrate(&model.User{})
	// database.Database.AutoMigrate(&model.Entry{})
}

// func teardown() {
// migrator := initializers.DB.Migrator()
// migrator.DropTable(&models.Post{})
// }
func main() {
	r := gin.Default()
	// handling any non existing route and saving IP
	r.NoRoute(controllers.NoRoute)
	// healthcheck
	r.GET("/healthz", controllers.Healthcheck)
	// normal CRUD operations for IP's
	r.GET("/", controllers.PostCreate)
	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.PostIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)
	r.Run() // listen and serve on 0.0.0.0:8080
}
