package graph

import (
	"github.com/fk/gqlplayground/db"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Db *db.Database
}
