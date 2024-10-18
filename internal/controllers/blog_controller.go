package controllers

import (
	"bloggingplatformapi/internal/models"
	"bloggingplatformapi/internal/services"
	"bloggingplatformapi/internal/utils"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type BlogController struct {
	Service services.BlogService
}

func NewBlogController(service services.BlogService) *BlogController {
	return &BlogController{service}
}

// CreateBlog handles POST /blogs
func (c *BlogController) CreateBlog(ctx *gin.Context) {
	var blog models.Blog
	if err := ctx.ShouldBindJSON(&blog); err != nil {
		log.Errorf("error binding json: %v", err)
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate the blog
	if err := utils.ValidateBlog(&blog); err != nil {
		log.Errorf("error validating blog: %v", err)
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.Service.CreateBlog(&blog); err != nil {
		log.Errorf("error creating blog: %v", err)
		utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to create blog")
		return
	}

	utils.RespondWithJSON(ctx, http.StatusCreated, blog)
}

// GetBlog handles GET /blogs/:id
func (c *BlogController) GetBlog(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Errorf("error converting paramater to number: %v", err)
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid post ID")
		return
	}

	blog, err := c.Service.GetBlogByID(id)
	if err != nil {
		log.Errorf("error getting blog: %v", err)
		if err == sql.ErrNoRows {
			utils.RespondWithError(ctx, http.StatusNotFound, "Blog not found")
		} else {
			utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to retrieve blog")
		}
		return
	}

	utils.RespondWithJSON(ctx, http.StatusOK, blog)
}

// GetAllBlogs handles GET /blogs
func (c *BlogController) GetAllBlogs(ctx *gin.Context) {
	term := ctx.Query("term")
	blogs, err := c.Service.GetAllBlogs(term)
	if err != nil {
		log.Errorf("error getting blogs: %v", err)
		utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to retrieve blogs")
		return
	}

	utils.RespondWithJSON(ctx, http.StatusOK, blogs)
}

// UpdateBlog handles PUT /blogs/:id
func (c *BlogController) UpdateBlog(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Errorf("error updating blog: %v", err)
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid post ID")
		return
	}

	var blog models.Blog
	if err := ctx.ShouldBindJSON(&blog); err != nil {
		log.Errorf("error binding JSON: %v", err)
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate the blog
	if err := utils.ValidateBlog(&blog); err != nil {
		log.Errorf("error validating blog: %v", err)
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	blog.ID = id
	if err := c.Service.UpdateBlog(&blog); err != nil {
		log.Errorf("error updating blogs: %v", err)
		if err == sql.ErrNoRows {
			utils.RespondWithError(ctx, http.StatusNotFound, "Blog not found")
		} else {
			utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to update post")
		}
		return
	}

	updatedBlog, err := c.Service.GetBlogByID(id)
	if err != nil {
		log.Errorf("error getting blog: %v", err)
		utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to retrieve updated post")
		return
	}

	utils.RespondWithJSON(ctx, http.StatusOK, updatedBlog)
}

// DeleteBlog handles DELETE /blogs/:id
func (c *BlogController) DeleteBlog(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Errorf("error converting parameter to integer: %v", err)
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid blog ID")
		return
	}

	if err := c.Service.DeleteBlog(id); err != nil {
		log.Errorf("error deleting blog: %v", err)
		if err == sql.ErrNoRows {
			utils.RespondWithError(ctx, http.StatusNotFound, "Blog not found")
		} else {
			utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to delete blog")
		}
		return
	}

	ctx.Status(http.StatusNoContent)
}
