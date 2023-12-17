package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/jrab66/go-poc/initializers"
	"github.com/jrab66/go-poc/models"
)

func Healthcheck(c *gin.Context) {

	// Respond 200
	c.Status(200)

}
func NoRoute(c *gin.Context) {
	// get data off request body
	var body struct {
		Ip string
	}
	c.Bind(&body)
	// create request
	post := models.Post{Ip: ReverseIP(GetClientIP(c))}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	//
	c.String(http.StatusNotFound, "Route not found, but saving IP \n")
	c.JSON(http.StatusNotFound, gin.H{
		"post": post,
	})
}

func PostCreate(c *gin.Context) {
	// get data off request body
	var body struct {
		Ip string
	}
	c.Bind(&body)
	// create request
	post := models.Post{Ip: ReverseIP(GetClientIP(c))}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	//
	c.JSON(http.StatusCreated, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	// get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	//
	//respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostsShow(c *gin.Context) {
	// get id from URL
	id := c.Param("id")
	// get the posts
	var post models.Post
	initializers.DB.Find(&post, id)

	//respond with them
	c.JSON(200, gin.H{
		"post": post,
	})

}

// trying testing
//
//	func TestHelloEmpty(t *testing.T) {
//		msg, err := Hello("")
//		if msg != "" || err == nil {
//				t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
//		}
//	}
func IsSuperAnimal(animal string) bool {
	return strings.ToLower(animal) == "gopher"
}

//

func PostsUpdate(c *gin.Context) {
	// get id from URL
	id := c.Param("id")
	// get data off request body
	var body struct {
		Ip string
	}
	c.Bind(&body)
	// find the post that we are update it
	var post models.Post
	initializers.DB.Find(&post, id)
	// Update post
	initializers.DB.Model(&post).Updates(models.Post{Ip: body.Ip})

	//respond with them
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsDelete(c *gin.Context) {
	// Get id from URL
	id := c.Param("id")

	// Delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.Status(200)

}
