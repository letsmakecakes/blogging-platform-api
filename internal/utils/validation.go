package utils

import (
	"bloggingplatformapi/internal/models"
	"errors"
	"strings"
)

// ValidateBlog ensures that all required fields in Blog are populated.
func ValidateBlog(blog *models.Blog) error {
	if isEmpty(blog.Title) {
		return errors.New("title is required")
	}
	if isEmpty(blog.Content) {
		return errors.New("content is required")
	}
	if isEmpty(blog.Category) {
		return errors.New("category is required")
	}
	if len(blog.Tags) == 0 {
		return errors.New("at least one tag is required")
	}
	return nil
}

// isEmpty checks if a string is empty or consists solely of whitespace.
func isEmpty(value string) bool {
	return strings.TrimSpace(value) == ""
}
