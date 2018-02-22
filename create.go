package resource

import (
	"net/http"

	"github.com/manyminds/api2go"
)

// Create attempts to save a new model instance to the database.
func (r *Resource) Create(obj interface{}, req api2go.Request) (api2go.Responder, error) {
	if err := r.runGlobalHooks(Create, req); err != nil {
		return nil, err
	}
	r.runSetHooks(obj, req)
	if err := r.DB.Create(obj).Error; err != nil {
		return nil, err
	}
	return &api2go.Response{
		Res:  obj,
		Code: http.StatusCreated,
	}, nil
}
