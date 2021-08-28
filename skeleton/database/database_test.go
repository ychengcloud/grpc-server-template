package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptions(t *testing.T) {
	o := defaultOptions
	opts := []Option{
		WithDebug(true),
		WithMysqlUser("decker"),
		// WithMysqlPort(3396),
		WithSqliteName("test.db"),
	}
	for _, opt := range opts {
		opt(&o)
	}

	assert.Equal(t, o.debug, true)
	assert.Equal(t, o.mysqlOptions.user, "decker")
	assert.Equal(t, o.mysqlOptions.port, 3306)
	assert.Equal(t, o.sqliteOptions.name, "test.db")
}
