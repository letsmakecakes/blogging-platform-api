package utils

import (
	"bloggingplatformapi/internal/models"
	"errors"
	"strings"
)

func ValidateBlog(blog *models.Blog) error {
	if strings.TrimSpace(blog.Title) == "" {
		return errors.New("title is required")
	}
	if strings.TrimSpace(blog.Content) == "" {
		return errors.New("content is required")
	}
	if strings.TrimSpace(blog.Category) == "" {
		return errors.New("category is required")
	}
	if len(blog.Tags) == 0 {
		return errors.New("at least one tag is required")
	}
	return nil
}
