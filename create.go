package resource

import (
	"net/http"

	"github.com/manyminds/api2go"
)

// Create attempts to save a new model instance to the database.
func (r *Resource) Create(obj interface{}, req api2go.Request) (api2go.Responder, error) {
	p := &Params{
		Action:  BeforeCreate,
		Request: req,
		DB:      r.DB,
		Obj:     obj,
	}
	if err := r.runHooks(p); err != nil {
		return nil, err
	}
	if p.DB = p.DB.Create(p.Obj); p.DB.Error != nil {
		return nil, p.DB.Error
	}
	p.Action = AfterCreate
	if err := r.runHooks(p); err != nil {
		return nil, err
	}
	return &api2go.Response{
		Res:  p.Obj,
		Code: http.StatusCreated,
	}, nil
}
