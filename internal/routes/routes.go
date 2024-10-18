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
		blogs := api.Group("/blogs")
		{
			blogs.POST("/", blogController.CreateBlog)      // Create a new blog
			blogs.GET("/", blogController.GetAllBlogs)      // Get all blogs
			blogs.GET("/:id", blogController.GetBlog)       // Get a blog by ID
			blogs.PUT("/:id", blogController.UpdateBlog)    // Update a blog by ID
			blogs.DELETE("/:id", blogController.DeleteBlog) // Delete a blog by ID
		}
	}
}
