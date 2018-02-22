package resource

import (
	"github.com/jinzhu/gorm"
)

// Resource implements the interfaces necessary to use a GORM model with the
// api2go package.
type Resource struct {
	// DB is a pointer to an open database connection.
	DB *gorm.DB

	// Type is an instance of the model for this resource.
	Type interface{}

	// GlobalHooks is a list of callbacks to run before all actions.
	GlobalHooks []GlobalHook

	// GetHooks is a list of callbacks run before every action except create.
	GetHooks []GetHook

	// SetHooks is a list of callbacks run before create and update.
	SetHooks []SetHook

	// Fields is a list of valid field names for filtering.
	Fields []string
}
