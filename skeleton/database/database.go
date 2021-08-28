package database

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	ErrNoDatabaseDsn = errors.New("no database dsn")
)

func UpdateID(db *gorm.DB) {
	field := db.Statement.Schema.PrioritizedPrimaryField
	if field != nil {
		fmt.Println("field:", field)
		field.Set(db.Statement.ReflectValue, uuid.New())
	}
}

// NewGorm
func NewGorm(ctx context.Context, opts ...Option) (db *gorm.DB, err error) {
	o := defaultOptions
	for _, opt := range opts {
		opt(&o)
	}

	var dsn string
	var dialector gorm.Dialector
	switch o.dialect {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", o.mysqlOptions.user, o.mysqlOptions.password, o.mysqlOptions.host, o.mysqlOptions.port, o.mysqlOptions.name, o.mysqlOptions.charset)
		fmt.Println("dsn:", dsn)
		dialector = mysql.Open(dsn)
	case "sqlite":
		dialector = sqlite.Open(o.sqliteOptions.name)
	default:
		return nil, errors.Wrap(err, "dialect error")

	}

	db, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "gorm open database connection error")
	}

	if o.debug {
		db = db.Debug()
	}

	// gorm:create 之前
	// db.Callback().Create().Before("gorm:create").Register("update_id", UpdateID)

	if o.autoMigrate {
		// db.AutoMigrate(&models.Account{})
	}

	return db, err
}
