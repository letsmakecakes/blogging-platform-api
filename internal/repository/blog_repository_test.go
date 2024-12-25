package repository

import (
	"bloggingplatformapi/internal/models"
	"bloggingplatformapi/pkg/mock/dbmock"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// setupTest initializes the mock database repository for testing.
func setupTest(t *testing.T) (*sqlmock.Sqlmock, BlogRepository) {
	t.Helper()
	db, mock, err := dbmock.NewMockDB()
	assert.NoError(t, err)

	repo := NewBlogRepository(db)
	return &mock, repo
}

// mockTimeNow provides a fixed timestamp for consistent test results.
func mockTimeNow() time.Time {
	return time.Date(2024, 12, 25, 16, 0, 0, 0, time.UTC)
}

func TestBlogRepository_GetAll(t *testing.T) {
	t.Parallel()

	mock, repo := setupTest(t)
	defer (*mock).ExpectClose()

	blog := &models.Blog{
		Title:    "Test Title",
		Content:  "Test Content",
		Category: "Tech",
		Tags:     []string{"Go", "Testing"},
	}

	(*mock).ExpectQuery(`INSERT INTO blogs \(title, content, category, tags, created_at, updated_at\) VALUES \(\$1, \$2, \$3, \$4, NOW\(\), NOW\(\)\) RETURNING id, created_at, updated_at`).
		WithArgs(blog.Title, blog.Content, blog.Category, "Go,Testing").
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).
			AddRow(1, mockTimeNow(), mockTimeNow()))

	err := repo.Create(blog)
	assert.NoError(t, err)
	assert.Equal(t, 1, blog.ID)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestBlogRepository_GetByID(t *testing.T) {
	t.Parallel()

	mock, repo := setupTest(t)
	defer (*mock).ExpectClose()

	now := mockTimeNow()
	blogID := 1
	expectedBlog := &models.Blog{
		ID:        blogID,
		Title:     "Test Title",
		Content:   "Test Content",
		Category:  "Tech",
		Tags:      []string{"Go", "Testing"},
		CreatedAt: now,
		UpdatedAt: now,
	}

	(*mock).ExpectQuery(
		`SELECT id, title, content, category, tags, created_at, updated_at 
         FROM blogs WHERE id = \$1`,
	).WithArgs(blogID).WillReturnRows(sqlmock.NewRows([]string{"id", "title", "content", "category", "tags", "created_at", "updated_at"}).
		AddRow(expectedBlog.ID, expectedBlog.Title, expectedBlog.Content, expectedBlog.Category, "Go,Testing", now, now))

	blog, err := repo.GetByID(blogID)
	assert.NoError(t, err)
	assert.Equal(t, expectedBlog, blog)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestBlogRepository_Update(t *testing.T) {
	t.Parallel()

	mock, repo := setupTest(t)
	defer (*mock).ExpectClose()

	now := mockTimeNow()
	blog := &models.Blog{
		ID:       1,
		Title:    "Updated Title",
		Content:  "Updated Content",
		Category: "Tech",
		Tags:     []string{"Go", "GORM"},
	}

	(*mock).ExpectQuery(
		`UPDATE blogs 
         SET title = \$1, content = \$2, category = \$3, tags = \$4, updated_at = NOW\(\) 
         WHERE id = \$5 RETURNING updated_at`,
	).WithArgs(blog.Title, blog.Content, blog.Category, "Go,GORM", blog.ID).
		WillReturnRows(sqlmock.NewRows([]string{"updated_at"}).AddRow(now))

	err := repo.Update(blog)
	assert.NoError(t, err)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestBlogRepository_Delete(t *testing.T) {
	t.Parallel()

	mock, repo := setupTest(t)
	defer (*mock).ExpectClose()

	blogID := 1

	(*mock).ExpectExec(
		`DELETE FROM blogs WHERE id = \$1`,
	).WithArgs(blogID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Delete(blogID)
	assert.NoError(t, err)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}
