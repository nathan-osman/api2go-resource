package resource

import (
	"github.com/jinzhu/gorm"
	"github.com/manyminds/api2go"
)

// Action indicates the type of action being performed.
type Action int

const (
	// BeforeCreate indicates that a resource is about to be created.
	BeforeCreate Action = iota
	// AfterCreate indicates that a resource has been created.
	AfterCreate

	// BeforeDelete indicates that a resource is about to be destroyed.
	BeforeDelete
	// AfterDelete indicates that a resource has been destroyed.
	AfterDelete

	// BeforeUpdate indicates that a resource is about to be modified.
	BeforeUpdate
	// AfterUpdate indicates that a resource has been modified.
	AfterUpdate

	// BeforeFindAll indicates that multiple resources are about to be retrieved.
	BeforeFindAll
	// AfterFindAll indicates that multiple resources may have been retrieved.
	AfterFindAll

	// BeforeFindOne indicates that a single resource is about to be retrieved.
	BeforeFindOne
	// AfterFindOne indicates that a single resource may have been retrieved.
	AfterFindOne
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
