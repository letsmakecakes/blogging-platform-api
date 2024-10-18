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
			blogs.POST("/", blogController.CreateBlog)
			blogs.GET("/", blogController.GetAllBlogs)
			blogs.PUT("/:id", blogController.UpdateBlog)
			blogs.DELETE("/:id", blogController.DeleteBlog)
		}
	}
}
