package controllers

import (
	"bloggingplatformapi/internal/models"
	"bloggingplatformapi/internal/services"
	"bloggingplatformapi/internal/utils"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	Service services.BlogService
}

func NewBlogController(service services.BlogService) *BlogController {
	return &BlogController{service}
}

// CreatePost handles POST /posts
func (c *BlogController) CreateBlog(ctx *gin.Context) {
	var blog models.Blog
	if err := ctx.ShouldBindJSON(&blog); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate the post
	if err := utils.ValidateBlog(&blog); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.Service.CreateBlog(&blog); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to create blog")
		return
	}

	utils.RespondWithJSON(ctx, http.StatusCreated, blog)
}

// GetPost handles GET /posts/:id
func (c *BlogController) GetBlog(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid post ID")
		return
	}

	blog, err := c.Service.GetBlogByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.RespondWithError(ctx, http.StatusNotFound, "Blog not found")
		} else {
			utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to retrieve blog")
		}
		return
	}

	utils.RespondWithJSON(ctx, http.StatusOK, blog)
}

