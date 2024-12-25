package dbmock

import (
	"database/sql"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"time"
)

// NewMockDB initializes a new GORM DB instance with a mock database connection.
// It returns the GORM DB instance, the sql mock object for expectations, and an error if any.
func NewMockDB() (*sql.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	return sqlDB, mock, nil
}

// AnyTime is a custom matcher for sql mock to match any time.Time value.
type AnyTime struct {
}

// Match checks if the provided driver.Value is of type time.Time.
func (a AnyTime) Match(value driver.Value) bool {
	_, ok := value.(time.Time)
	return ok
}
