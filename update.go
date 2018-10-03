package resource

import (
	"net/http"

	"github.com/manyminds/api2go"
)

// Update attempts to update a model instance in the database.
func (r *Resource) Update(obj interface{}, req api2go.Request) (api2go.Responder, error) {
	p := &Params{
		Action:  Update,
		Request: req,
		DB:      r.DB,
		Obj:     obj,
	}
	if err := r.runHooks(p); err != nil {
		return nil, err
	}
	if err := translateError(p.DB.Model(r.Type).Updates(p.Obj)); err != nil {
		return nil, err
	}
	return &api2go.Response{
		Res:  p.Obj,
		Code: http.StatusOK,
	}, nil
}
