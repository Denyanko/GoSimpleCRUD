package main

import (
	"github.com/Denyanko/GoSimpleCRUD/controllers"
	"github.com/Denyanko/GoSimpleCRUD/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.POST("/api/v1/posts", controllers.CreatePost)
	router.GET("/api/v1/posts", controllers.FindPosts)
	router.GET("/api/v1/posts/:id", controllers.FindPost)
	router.PATCH("/api/v1/posts/:id", controllers.UpdatePost)
	router.DELETE("/api/v1/posts/:id", controllers.DeletePost)

	_ = router.Run("localhost:8080")
}
