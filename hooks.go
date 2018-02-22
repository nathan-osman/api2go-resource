package resource

import (
	"github.com/jinzhu/gorm"
	"github.com/manyminds/api2go"
)

type Action int

const (
	Create Action = iota
	Delete
	FindAll
	FindOne
	Update
)

// GlobalHook is a callback that is run immediately before processing any
// action. If the hook returns an error, the action does not run.
type GlobalHook func(Action, api2go.Request) error

// GetHook is a callback that is run before all actions except create. The
// database connection is passed as a parameter and the return value is then
// used by the action.
type GetHook func(*gorm.DB, api2go.Request) *gorm.DB

// SetHook is a callback that is run before the create and update actions. The
// hook may modify the object to prepare it for the database.
type SetHook func(interface{}, api2go.Request)

// runGlobalHooks executes all global hooks, returning an error if any hook
// fails.
func (r *Resource) runGlobalHooks(action Action, req api2go.Request) error {
	for _, h := range r.GlobalHooks {
		if err := h(action, req); err != nil {
			return err
		}
	}
	return nil
}

// runGetHooks executes all get hooks in order, returning the database
// connection returned by the last hook.
func (r *Resource) runGetHooks(c *gorm.DB, req api2go.Request) *gorm.DB {
	for _, h := range r.GetHooks {
		c = h(c, req)
	}
	return c
}

// runSetHooks executes all set hooks.
func (r *Resource) runSetHooks(obj interface{}, req api2go.Request) {
	for _, h := range r.SetHooks {
		h(obj, req)
	}
}
