# Project Documentation

Welcome to the **Blogging Platform API** documentation. This project provides RESTful API services for managing blogs, including CRUD operations and search functionality.

## Table of Contents

1. [API Overview](#api-overview)
2. [OpenAPI Specification](#openapi-specification)
3. [Using the API](#using-the-api)
4. [Examples](#examples)

---

## API Overview

The Blogging Platform API supports the following features:
- Creating, retrieving, updating, and deleting blogs.
- Searching blogs by title or content.
- Filtering by categories and tags.

The API is implemented using the Gin framework and follows best practices for RESTful API design.

---

## OpenAPI Specification

The API is documented using the OpenAPI 3.0 specification. You can find the complete spec in YAML format [here](openapi.yaml).

### Viewing the API Docs
To view the API documentation interactively:
1. Use a tool like Swagger UI, ReDoc, or Postman.
2. Load the `openapi.yaml` or `openapi.json` file.

---

## Using the API

### Base URL
http://localhost:<port>

javascript
Copy code

Replace `<port>` with the configured port for the server.

### Authentication (in Future)
All endpoints require JWT-based authentication. Include the token in the `Authorization` header:
Authorization: Bearer <your-token>

---

## Examples

Refer to the `examples/` directory for sample requests and responses.

- **Create Blog**: [examples/create_blog.json](examples/create_blog.json)
- **Update Blog**: [examples/update_blog.json](examples/update_blog.json)
- **Error Response**: [examples/error_response.json](examples/error_response.json)
