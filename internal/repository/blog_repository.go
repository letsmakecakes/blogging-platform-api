package repository

import (
	"bloggingplatformapi/internal/models"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"strings"
)

// BlogRepository defines the interfaces for blog-related database operations.
type BlogRepository interface {
	Create(blog *models.Blog) error             // Creates a new blog
	GetByID(id int) (*models.Blog, error)       // Fetch a blog by its ID
	GetAll(term string) ([]*models.Blog, error) // Fetch all blogs, optional filtered by a search term
	Update(blog *models.Blog) error             // Update an existing blog
	Delete(id int) error                        // Delete a blog by its ID
}

// blogRepository is a concrete implementation of the BlogRepository interface.
type blogRepository struct {
	db *sql.DB // Database connection
}

// NewBlogRepository creates a new BlogRepository instance.
func NewBlogRepository(db *sql.DB) BlogRepository {
	return &blogRepository{db}
}

// Create inserts a new blog into the database.
func (r *blogRepository) Create(blog *models.Blog) error {
	query := `
		INSERT INTO blogs (title, content, category, tags, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW()) 
		RETURNING id, created_at, updated_at
	`
	tags := strings.Join(blog.Tags, ",") // Convert tags slice into a comma-seperated string
	err := r.db.QueryRow(query, blog.Title, blog.Content, blog.Category, tags).Scan(&blog.ID, &blog.CreatedAt, &blog.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

// GetByID retrieves a single blog by its ID.
func (r *blogRepository) GetByID(id int) (*models.Blog, error) {
	query := `
		SELECT id, title, content, category, tags, created_at, updated_at 
		FROM blogs 
		WHERE id = $1
	`
	var blog models.Blog
	var tags string

	err := r.db.QueryRow(query, id).Scan(
		&blog.ID, &blog.Title, &blog.Content, &blog.Category, &tags, &blog.CreatedAt, &blog.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	blog.Tags = strings.Split(tags, ",") // Convert the tags string back to a slice
	return &blog, nil
}

// GetAll retrieves all blogs or filters them based on a search term.
func (r *blogRepository) GetAll(term string) ([]*models.Blog, error) {
	var blogs []*models.Blog
	var rows *sql.Rows
	var err error

	if term != "" {
		likeTerm := "%" + term + "%"
		query := `
			SELECT id, title, content, category, tags, created_at, updated_at
			From blogs
			WHERE title ILIKE $1 OR content ILIKE $1 OR category ILIKE $1
		`
		rows, err = r.db.Query(query, likeTerm)
	} else {
		query := `SELECT id, title, content, category, tags, created_at, updated_at FROM blogs`
		rows, err = r.db.Query(query)
	}

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Errorf("error closing rows: %v", err)
		}
	}(rows) // Ensure rows are properly closed

	for rows.Next() {
		var blog models.Blog
		var tags string
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Category, &tags, &blog.CreatedAt, &blog.UpdatedAt); err != nil {
			return nil, err
		}
		blog.Tags = strings.Split(tags, ",")
		blogs = append(blogs, &blog)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}

// Update modifies an existing blog in the database.
func (r *blogRepository) Update(blog *models.Blog) error {
	query := `
		UPDATE blogs
		SET title = $1, content = $2, category = $3, tags = $4, updated_at = NOW()
		WHERE id = $5 
		RETURNING updated_at
	`
	tags := strings.Join(blog.Tags, ",") // Convert tags slice to a comma-seperated string
	err := r.db.QueryRow(query, blog.Title, blog.Content, blog.Category, tags, blog.ID).Scan(&blog.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a blog by its ID from the database.
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
		return sql.ErrNoRows // Return a specific error if no rows were deleted
	}

	return nil
}
