package models

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()

	ErrAccountNotFound = errors.New("account not found")
	ErrEmailExists     = errors.New("email already exists")
	ErrNoDatabase      = errors.New("no database connection details")
	ErrNoPasswordGiven = errors.New("a password is required")
	ErrBadRequest      = errors.New("bad request")
)

type FilterLogicOp string

const (
	FilterLogicOpOr  FilterLogicOp = "OR"
	FilterLogicOpAnd FilterLogicOp = "AND"
)

var FilterLogicOpAll = map[string]string{
	"Eq":         "= ?",
	"Not":        "<> = ?",
	"In":         "IN (?)",
	"Gt":         "> ?",
	"Gte":        ">= ?",
	"Lt":         "< ?",
	"Lte":        "<= ?",
	"Contains":   "LIKE (?)",
	"StartsWith": "LIKE (?%)",
	"EndsWith":   "LIKE (%?)",
}
