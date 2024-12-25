package services

import (
	"bloggingplatformapi/internal/models"
	"bloggingplatformapi/internal/repository"
)

// BlogService defines the contract for blog-related operations.
type BlogService interface {
	CreateBlog(blog *models.Blog) error
	GetBlogByID(id int) (*models.Blog, error)
	GetAllBlogs(term string) ([]*models.Blog, error)
	UpdateBlog(blog *models.Blog) error
	DeleteBlog(id int) error
}

// blogService implements the BlogService interface.
type blogService struct {
	repo repository.BlogRepository
}

// NewBlogService creates a new instance of BlogService with the provided repository.
func NewBlogService(repo repository.BlogRepository) BlogService {
	return &blogService{repo}
}

// CreateBlog delegates the creation of a blog to the repository layer.
func (s *blogService) CreateBlog(blog *models.Blog) error {
	return s.repo.Create(blog)
}

// GetBlogByID retrieves a single blog by its ID from the repository layer.
func (s *blogService) GetBlogByID(id int) (*models.Blog, error) {
	return s.repo.GetByID(id)
}

// GetAllBlogs retrieves all blogs matching the search term from the repository.
func (s *blogService) GetAllBlogs(term string) ([]*models.Blog, error) {
	return s.repo.GetAll(term)
}

// UpdateBlog updates an existing blog via the repository layer.
func (s *blogService) UpdateBlog(blog *models.Blog) error {
	return s.repo.Update(blog)
}

// DeleteBlog removes a blog by its ID using the repository layer.
func (s *blogService) DeleteBlog(id int) error {
	return s.repo.Delete(id)
}
