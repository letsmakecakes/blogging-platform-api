package controllers

import (
	"bloggingplatformapi/internal/models"
	"bloggingplatformapi/internal/services"
	"bloggingplatformapi/internal/utils"
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// BlogController is responsible for handling HTTP requests related to blogs.
type BlogController struct {
	Service services.BlogService
}

// NewBlogController creates a new instance of BlogController with the provided BlogService.
func NewBlogController(service services.BlogService) *BlogController {
	return &BlogController{service}
}

// CreateBlog handles the creation of a new blog via POST /blogs.
// It validates the incoming request, processes the creation through the service layer, and returns the created blog.
func (c *BlogController) CreateBlog(ctx *gin.Context) {
	var blog models.Blog
	// Bind incoming JSON payload to the Blog model.
	if err := ctx.ShouldBindJSON(&blog); err != nil {
		logAndRespond(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Validate the blog content.
	if err := utils.ValidateBlog(&blog); err != nil {
		logAndRespond(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	// Pass the blog to the service for creation.
	if err := c.Service.CreateBlog(&blog); err != nil {
		logAndRespond(ctx, http.StatusInternalServerError, "Failed to create blog", err)
		return
	}

	utils.RespondWithJSON(ctx, http.StatusCreated, blog)
}

// GetBlog retrieves a specific key by its ID via GET /blogs/:id.
// It validates the ID parameter and fetches the blog from the service layer.
func (c *BlogController) GetBlog(ctx *gin.Context) {
	id, err := parseID(ctx.Param("blogId"))
	if err != nil {
		logAndRespond(ctx, http.StatusBadRequest, "Invalid blog ID", err)
		return
	}

	// Fetch the blog from the service layer.
	blog, err := c.Service.GetBlogByID(id)
	if handleServiceError(ctx, err, "Failed to retrieve blog") {
		return
	}

	utils.RespondWithJSON(ctx, http.StatusOK, blog)
}

// GetAllBlogs retrieves all blogs or filters them based on a search term via GET /blogs.
// It handles optional query parameters and fetches the blogs from the service layer.
func (c *BlogController) GetAllBlogs(ctx *gin.Context) {
	term := ctx.Query("term")
	// Fetch blogs with an optional search term.
	blogs, err := c.Service.GetAllBlogs(term)
	if err != nil {
		logAndRespond(ctx, http.StatusInternalServerError, "Failed to retrieve blogs", err)
		return
	}

	utils.RespondWithJSON(ctx, http.StatusOK, blogs)
}

// UpdateBlog updates an existing blog post by its ID via PUT /blogs/:id.
// It validates the ID parameter, incoming payload, and performs the update through the service layer.
func (c *BlogController) UpdateBlog(ctx *gin.Context) {
	id, err := parseID(ctx.Param("blogId"))
	if err != nil {
		logAndRespond(ctx, http.StatusBadRequest, "Invalid blog ID", err)
		return
	}

	var blog models.Blog
	// Bind incoming JSON payload to the Blog model.
	if err := ctx.ShouldBindJSON(&blog); err != nil {
		logAndRespond(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Validate the blog content.
	if err := utils.ValidateBlog(&blog); err != nil {
		logAndRespond(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	// Assign the blog ID and pass it to the service for update.
	blog.ID = id
	if err := c.Service.UpdateBlog(&blog); err != nil {
		if handleServiceError(ctx, err, "Failed to update blog") {
			return
		}
	}

	// Fetch the updated blog to ensure successful update.
	updatedBlog, err := c.Service.GetBlogByID(id)
	if handleServiceError(ctx, err, "Failed to retrieve updated blog") {
		return
	}

	utils.RespondWithJSON(ctx, http.StatusOK, updatedBlog)
}

// DeleteBlog handles DELETE /blogs/:id
func (c *BlogController) DeleteBlog(ctx *gin.Context) {
	id, err := parseID(ctx.Param("blogId"))
	if err != nil {
		logAndRespond(ctx, http.StatusBadRequest, "Invalid blog ID", err)
		return
	}

	// Perform the deletion through the service layer.
	if err := c.Service.DeleteBlog(id); err != nil {
		if handleServiceError(ctx, err, "Failed to delete blog") {
			return
		}
	}

	ctx.Status(http.StatusNoContent)
}

// Helper functions

// parseID converts a string parameter to an integer.
// It returns an error if the conversion fails.
func parseID(param string) (int, error) {
	return strconv.Atoi(param)
}

// logAndRespond logs the error and sends a JSON response with the provided status code and message.
func logAndRespond(ctx *gin.Context, status int, message string, err error) {
	log.Errorf("%s: %v", message, err)
	utils.RespondWithError(ctx, status, message)
}

// handleServiceError handles common errors from the service layer.
// It checks for specific conditions (e.g., sql.ErrNoRows) and responds accordingly.
// Returns true if an error is handler, otherwise false.
func handleServiceError(ctx *gin.Context, err error, message string) bool {
	if err != nil {
		log.Errorf("%s: %v", message, err)
		if errors.Is(err, sql.ErrNoRows) {
			utils.RespondWithError(ctx, http.StatusNotFound, "Blog not found")
		} else {
			utils.RespondWithError(ctx, http.StatusInternalServerError, message)
		}
		return true
	}
	return false
}
