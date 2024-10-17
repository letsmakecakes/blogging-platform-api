package controllers

import (
	"bloggingplatformapi/internal/models"
	"bloggingplatformapi/internal/services"
	"bloggingplatformapi/internal/utils"
	"net/http"

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
