package repository

import (
	"bloggingplatformapi/internal/models"
	"database/sql"
	"strings"
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

func (r *blogRepository) Create(blog *models.Blog) error {
	query := `INSERT INTO posts (title, content, category, tags, created_at, updated_at)
				VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id, created_at, updated_at`
	tags := strings.Join(blog.Tags, ",")
	err := r.db.QueryRow(query, blog.Title, blog.Content, blog.Category, tags).Scan(&blog.ID, &blog.CreatedAt, &blog.UpdatedAt)
	return err
}
