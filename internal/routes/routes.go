package routes

import (
	"bloggingplatformapi/internal/controllers"
	"bloggingplatformapi/internal/repository"
	"bloggingplatformapi/internal/services"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	// Initialize repository, service, and controller
	blogRepo := repository.NewBlogRepository(db)
	blogService := services.NewBlogService(blogRepo)
	blogController := controllers.NewBlogController(blogService)

	// Group routes under /api for versioning or future scalability
	api := router.Group("/api")
	{
		posts := api.Group("/blogs") // Renamed to 'posts'
		{
			posts.POST("/", blogController.CreateBlog)      // Create a new post
			posts.GET("/", blogController.GetAllBlogs)      // Get all posts
			posts.GET("/:id", blogController.GetBlog)       // Get a post by ID
			posts.PUT("/:id", blogController.UpdateBlog)    // Update a post by ID
			posts.DELETE("/:id", blogController.DeleteBlog) // Delete a post by ID
		}
	}
}
