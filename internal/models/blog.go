package models

import "time"

// Blog represents a blog post with its metadata, content, and categorization.
type Blog struct {
	ID        int       `json:"id"`                          // Unique identifier for the blog
	Title     string    `json:"title" binding:"required"`    // Title of the blog (required)
	Content   string    `json:"content" binding:"required"`  // Content of the blog (required)
	Category  string    `json:"category" binding:"required"` // Blog category (required)
	Tags      []string  `json:"tags" binding:"required"`     // Tags associated with the blog (required)
	CreatedAt time.Time `json:"createdAt"`                   // Timestamp when the blog was created
	UpdatedAt time.Time `json:"updatedAt"`                   // Timestamp when the blog was last updated
}
