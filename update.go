package resource

import (
	"net/http"

	"github.com/manyminds/api2go"
)

// Update attempts to update a model instance in the database.
func (r *Resource) Update(obj interface{}, req api2go.Request) (api2go.Responder, error) {
	p := &Params{
		Action:  BeforeUpdate,
		Request: req,
		DB:      r.DB,
		Obj:     obj,
	}
	if err := r.runHooks(p); err != nil {
		return nil, err
	}
	if p.DB = p.DB.Model(r.Type).Updates(p.Obj); p.DB.Error != nil {
		return nil, p.DB.Error
	}
	p.Action = AfterUpdate
	if err := r.runHooks(p); err != nil {
		return nil, err
	}
	return &api2go.Response{
		Res:  p.Obj,
		Code: http.StatusOK,
	}, nil
}
