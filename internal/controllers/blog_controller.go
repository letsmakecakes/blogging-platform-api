package controllers

import "bloggingplatformapi/internal/services"

type BlogController struct {
	Service services.BlogService
}

func NewBlogController(service services.BlogService) *BlogController {
	return &BlogController{service}
}
