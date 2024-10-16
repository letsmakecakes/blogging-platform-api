package services

import (
	"bloggingplatformapi/internal/models"
	"bloggingplatformapi/internal/repository"
)

type BlogService interface {
	CreateBlog(blog *models.Blog) error
	GetBlogByID(id int) (*models.Blog, error)
	GetAllBlogs(term string) ([]*models.Blog, error)
	UpdateBlog(blog *models.Blog) error
	DeleteBlog(id int) error 
}

type blogService struct {
	repo repository.BlogRepository
}
