package repository

import (
	"bloggingplatformapi/internal/models"
	"database/sql"
)

type BlogRepository interface {
	Create(post *models.Blog) error
	GetByID(id int) (*models.Blog, error)
	GetAll(term string) ([]*models.Blog, error)
	Update(post *models.Blog) error
	Delete(id int) error
}

type blogRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) BlogRepository {
	return &blogRepository{db}
}
