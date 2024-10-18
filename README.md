# Blogging Platform API

A RESTful API for a blogging platform built using the [Gin Web Framework](https://github.com/gin-gonic/gin) in Golang. The API supports CRUD operations for blog posts, including optional filtering based on search terms.

## Table of Contents

- [Features](#features)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Environment Variables](#environment-variables)
- [Running the Project](#running-the-project)
- [API Endpoints](#api-endpoints)
- [Database Migrations](#database-migrations)
- [Contributing](#contributing)
- [License](#license)

## Features

- Create, read, update, and delete (CRUD) blog posts
- Filter blog posts based on title, content, or category
- API built with [Gin](https://github.com/gin-gonic/gin) framework
- PostgreSQL integration
- Struct-based validation
- Easily extendable and customizable architecture
- Structured logging
- Environment-based configuration

## Project Structure

```
blogging-platform-api/
├── cmd/
│   └── server/
│       └── main.go         # Entry point for the application
├── internal/
│   ├── config/
│   │   └── config.go       # Configuration loader
│   ├── controllers/
│   │   └── blog_controller.go # API Handlers for blog posts
│   ├── models/
│   │   └── blog.go         # Blog post model
│   ├── repository/
│   │   └── blog_repository.go # Database operations for blog posts
│   ├── routes/
│   │   └── routes.go       # Route definitions
│   ├── services/
│   │   └── blog_service.go # Business logic for blog posts
│   └── utils/
│       ├── response.go     # Utility functions for API responses
│       └── validation.go   # Validation utilities
├── pkg/
│   └── db/
│       └── db.go           # Database connection and initialization
├── migrations/
│   └── 001_create_blogs_table.sql # SQL migration for blogs table
├── .env                     # Environment variables
├── .gitignore                # Git ignore file
├── go.mod                    # Go module definition
├── go.sum                    # Go module dependencies
└── README.md                 # Project documentation
```

## Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/your-username/blogging-platform-api.git
   cd blogging-platform-api
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Create and configure your `.env` file. (See [Environment Variables](#environment-variables) below.)

4. Run database migrations (PostgreSQL) to create the necessary tables:

   ```bash
   psql -h <host> -d <database> -U <user> -f migrations/001_create_posts_table.sql
   ```

## Environment Variables

To run this project, you'll need to set up the following environment variables in your `.env` file:

```bash
# Application
PORT=8080
ENVIRONMENT=development  # or 'production'

# Database (PostgreSQL)
DATABASE_URL=postgres://<username>:<password>@<host>:<port>/<database>?sslmode=disable
```

## Running the Project

To run the application locally, use the following command:

```bash
go run cmd/server/main.go
```

The API will be accessible at `http://localhost:8080`.

## API Endpoints

### Blog Posts

- **GET** `/blogs/`: Fetch all blog posts. Supports filtering via query parameters (e.g., `term`).
- **GET** `/blogs/:id`: Fetch a single blog post by ID.
- **POST** `/blogs`: Create a new blog post. Requires JSON payload.
- **PUT** `/blogs/:id`: Update an existing blog post by ID.
- **DELETE** `/blogs/:id`: Delete a blog post by ID.

#### Example Request and Response

##### Create a New Post

**Request:**

```bash
POST /blogs
Content-Type: application/json
```

```json
{
  "title": "My First Blog Post",
  "content": "This is the content of my first post!",
  "category": "Tech",
  "tags": ["Go", "Programming", "Backend"]
}
```

**Response:**

```json
{
  "id": 1,
  "title": "My First Blog Post",
  "content": "This is the content of my first post!",
  "category": "Tech",
  "tags": ["Go", "Programming", "Backend"],
  "createdAt": "2024-10-16T14:45:00Z",
  "updatedAt": "2024-10-16T14:45:00Z"
}
```

##### Get All Posts

**Request:**

```bash
GET /blogs?term=Tech
```

**Response:**

```json
[
  {
    "id": 1,
    "title": "My First Blog Post",
    "content": "This is the content of my first post!",
    "category": "Tech",
    "tags": ["Go", "Programming", "Backend"],
    "createdAt": "2024-10-16T14:45:00Z",
    "updatedAt": "2024-10-16T14:45:00Z"
  }
]
```

## Database Migrations

Ensure you run the SQL migration file located in the `migrations/` directory to create the `posts` table.

Run the migration file with:

```bash
psql -h <host> -d <database> -U <user> -f migrations/001_create_posts_table.sql
```

The migration file `001_create_posts_table.sql` contains the SQL necessary to create the `posts` table:

```sql
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    category VARCHAR(100) NOT NULL,
    tags TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
```

## Contributing

If you'd like to contribute to this project, please fork the repository and submit a pull request. Any improvements and suggestions are welcome!

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/new-feature`)
3. Commit your changes (`git commit -m 'Add some feature'`)
4. Push to the branch (`git push origin feature/new-feature`)
5. Open a pull request

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
