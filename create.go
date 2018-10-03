package resource

import (
	"net/http"

	"github.com/manyminds/api2go"
)

// Create attempts to save a new model instance to the database.
func (r *Resource) Create(obj interface{}, req api2go.Request) (api2go.Responder, error) {
	p := &Params{
		Action:  Create,
		Request: req,
		DB:      r.DB,
		Obj:     obj,
	}
	if err := r.runHooks(p); err != nil {
		return nil, err
	}
	if err := p.DB.Create(p.Obj).Error; err != nil {
		return nil, err
	}
	return &api2go.Response{
		Res:  p.Obj,
		Code: http.StatusCreated,
	}, nil
}
