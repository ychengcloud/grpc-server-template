package api

import (
	"embed"
)

//go:embed openapi/*
var OpenAPI embed.FS
