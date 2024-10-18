package repository

import (
	"bloggingplatformapi/internal/models"
	"database/sql"
	"strings"
)

type BlogRepository interface {
	Create(blog *models.Blog) error
	GetByID(id int) (*models.Blog, error)
	GetAll(term string) ([]*models.Blog, error)
	Update(blog *models.Blog) error
	Delete(id int) error
}

type blogRepository struct {
	db *sql.DB
}

func NewBlogRepository(db *sql.DB) BlogRepository {
	return &blogRepository{db}
}

func (r *blogRepository) Create(blog *models.Blog) error {
	query := `INSERT INTO posts (title, content, category, tags, created_at, updated_at)
				VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id, created_at, updated_at`
	tags := strings.Join(blog.Tags, ",")
	err := r.db.QueryRow(query, blog.Title, blog.Content, blog.Category, tags).Scan(&blog.ID, &blog.CreatedAt, &blog.UpdatedAt)
	return err
}

func (r *blogRepository) GetByID(id int) (*models.Blog, error) {
	query := `SELECT id, title, content, category, tags, created_at, updated_at FROM posts WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var blog models.Blog
	var tags string
	err := row.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Category, &tags, &blog.CreatedAt, &blog.UpdatedAt)
	if err != nil {
		return nil, err
	}

	blog.Tags = strings.Split(tags, ",")
	return &blog, nil
}

func (r *blogRepository) GetAll(term string) ([]*models.Blog, error) {
	var blogs []*models.Blog
	var rows *sql.Rows
	var err error

	if term != "" {
		likeTerm := "%" + term + "%"
		query := `SELECT id, title, content, category, tags, created_at, updated_at
					From posts
					WHERE title ILIKE $1 OR content ILIKE $1 OR category ILIKE $1`
		rows, err = r.db.Query(query, likeTerm)
	} else {
		query := `SELECT id, title, content, category, tags, created_at, updated_at FROM posts`
		rows, err = r.db.Query(query)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var blog models.Blog
		var tags string
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Category, &tags, &blog.CreatedAt, &blog.UpdatedAt); err != nil {
			return nil, err
		}
		blog.Tags = strings.Split(tags, ",")
		blogs = append(blogs, &blog)
	}

	return blogs, nil
}

func (r *blogRepository) Update(blog *models.Blog) error {
	query := `UPDATE posts SET title = $1, content = $2, category = $3, tags = $4, updated_at = NOW()
				WHERE id = $5 RETURNING updated_at`
	tags := strings.Join(blog.Tags, ",")
	err := r.db.QueryRow(query, blog.Title, blog.Content, blog.Category, tags, blog.ID).Scan(&blog.UpdatedAt)
	return err
}

func (r *blogRepository) Delete(id int) error {
	query := `DELETE FROM blogs WHERE id = $1`
	res, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
