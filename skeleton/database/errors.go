package database

import "strings"

// IsUniqueConstraintError reports if the error resulted from a DB uniqueness constraint violation.
// e.g. duplicate value in unique index.
func IsUniqueConstraintError(err error) bool {
	uniquenessErrors := []string{
		"Error 1062",                 // MySQL
		"violates unique constraint", // Postgres
		"UNIQUE constraint failed",   // SQLite
	}
	for _, s := range uniquenessErrors {
		if strings.Contains(err.Error(), s) {
			return true
		}
	}
	return false
}

// IsForeignKeyConstraintError reports if the error resulted from a DB FK constraint violation.
// e.g. parent row does not exist.
func IsForeignKeyConstraintError(err error) bool {
	fkErrors := []string{
		"Error 1452",                      // MySQL
		"violates foreign key constraint", // Postgres
		"FOREIGN KEY constraint failed",   // SQLite
	}
	for _, s := range fkErrors {
		if strings.Contains(err.Error(), s) {
			return true
		}
	}
	return false
}
