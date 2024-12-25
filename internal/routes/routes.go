package routes

import (
	"bloggingplatformapi/internal/controllers"
	"bloggingplatformapi/internal/repository"
	"bloggingplatformapi/internal/services"
	"database/sql"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes all application routes.
func SetupRoutes(router *gin.Engine, db *sql.DB) {
	// Setup blog module dependencies
	blogController := initializeBlogController(db)

	// Define API routes
	api := router.Group("/api")
	setupBlogRoutes(api, blogController)
}

// initializeBlogController sets up the blog controller with its dependencies.
func initializeBlogController(db *sql.DB) *controllers.BlogController {
	blogRepo := repository.NewBlogRepository(db)      // Initialize the repository
	blogService := services.NewBlogService(blogRepo)  // Initialize the service
	return controllers.NewBlogController(blogService) // Initialize the controller
}

// setupBlogRoutes configures routes for blog-related operations.
func setupBlogRoutes(api *gin.RouterGroup, blogController *controllers.BlogController) {
	blogs := api.Group("/blogs")
	{
		blogs.POST("", blogController.CreateBlog)       // Create a new blog
		blogs.GET("", blogController.GetAllBlogs)       // Get all blogs
		blogs.GET("/:id", blogController.GetBlog)       // Get a blog by ID
		blogs.PUT("/:id", blogController.UpdateBlog)    // Update a blog by ID
		blogs.DELETE("/:id", blogController.DeleteBlog) // Delete a blog by ID
	}
}
