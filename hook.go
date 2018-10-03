package resource

import (
	"github.com/jinzhu/gorm"
	"github.com/manyminds/api2go"
)

// Action indicates the type of action being performed.
type Action int

const (
	// Create indicates that a resource is being created.
	Create Action = iota
	// Delete indicates that a resource is being destroyed.
	Delete
	// FindAll attempts to filter a set of resources.
	FindAll
	// FindOne attempts to find a single resource by its primary key.
	FindOne
	// Update indicates that a resource is being updated.
	Update
)

// Params contains information about an API request. Only certain members will
// contain valid data, depending on the action.
type Params struct {
	Action  Action
	Request api2go.Request
	DB      *gorm.DB
	Obj     interface{}
}

// Hook is a callback that is run immediately before processing each action. If
// the hook returns an error, the action does not run.
type Hook func(*Params) error

// runHooks executes all hooks for a resource. If the return value is non-nil,
// processing should stop. The parameter values may be modified to affect the
// operation.
func (r *Resource) runHooks(params *Params) error {
	for _, h := range r.Hooks {
		if err := h(params); err != nil {
			return err
		}
	}
	return nil
}
