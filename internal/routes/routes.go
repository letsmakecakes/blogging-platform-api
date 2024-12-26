package routes

import (
	"bloggingplatformapi/internal/controllers"
	"bloggingplatformapi/internal/repository"
	"bloggingplatformapi/internal/services"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

// SetupRoutes initializes all application routes.
func SetupRoutes(router *gin.Engine, db *sql.DB) {
	// Apply CORS middleware
	router.Use(createCORSHandler())

	// Setup blog module dependencies
	blogController := initializeBlogController(db)

	// Define API routes
	api := router.Group("/api/v1")
	setupBlogRoutes(api, blogController)
}

// createCORSHandler creates and returns a Gin-compatible CORS middleware.
func createCORSHandler() gin.HandlerFunc {
	corsMiddleware := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	})
	return func(c *gin.Context) {
		corsMiddleware.HandlerFunc(c.Writer, c.Request)
		c.Next()
	}
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
		blogs.GET("", blogController.GetAllBlogs)           // List all blogs
		blogs.POST("", blogController.CreateBlog)           // Create a new blog
		blogs.GET("/:blogId", blogController.GetBlog)       // Get a specific blog
		blogs.PUT("/:blogId", blogController.UpdateBlog)    // Update a specific blog
		blogs.DELETE("/:blogId", blogController.DeleteBlog) // Delete a specific blog
	}

	// Additional routes for future scalability (example: tags)
	//api.GET("/tags", blogController.GetAllTags)                   // List all tags
	//api.GET("/tags/:tagName/blogs", blogController.GetBlogsByTag) // Get all blogs for a specific tag

	// Example author routes for extensibility
	//api.GET("/authors", blogController.GetAllAuthors)                    // List all authors
	//api.GET("/authors/:authorId", blogController.GetAuthorDetails)       // Get author details
	//api.GET("/authors/:authorId/blogs", blogController.GetBlogsByAuthor) // Get all blogs by an author
}
