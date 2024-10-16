package repository

import "bloggingplatformapi/internal/models"

type BlogRepository interface {
	Create(post *models.Blog) error
	GetByID(id int) (*models.Blog, error)
	GetAll(term string) ([]*models.Blog, error)
	Update(post *models.Blog) error
	Delete(id int) error 
}
