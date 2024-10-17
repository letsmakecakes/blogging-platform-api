package utils

import (
	"bloggingplatformapi/internal/models"
	"errors"
	"strings"
)

func ValidateBlog(blog *models.Blog) error {
	if strings.TrimSpace(blog.Title) == "" {
		return errors.New("Title is required")
	}
	if strings.TrimSpace(blog.Content) == "" {
		return errors.New("Content is required")
	}
	if strings.TrimSpace(blog.Category) == "" {
		return errors.New("Category is required")
	}
	if len(blog.Tags) == 0 {
		return errors.New("At least one tag is required")
	}
	return nil
}
