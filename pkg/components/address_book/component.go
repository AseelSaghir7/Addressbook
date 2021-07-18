package address_book

import (
	"database/sql"
)

// Component users is for managing users with httprouter.Handle(s) and validators
type Component struct {
	db          *sql.DB
}

// New create new users component
func New(db *sql.DB) Component {
	return Component{
		db:          db,
	}
}
