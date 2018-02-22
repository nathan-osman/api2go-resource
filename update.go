package resource

import (
	"net/http"

	"github.com/manyminds/api2go"
)

// Update attempts to update a model instance in the database.
func (r *Resource) Update(obj interface{}, req api2go.Request) (api2go.Responder, error) {
	if err := r.runGlobalHooks(Update, req); err != nil {
		return nil, err
	}
	c := r.runGetHooks(r.DB, req)
	r.runSetHooks(obj, req)
	if err := c.Model(r.Type).Updates(obj).Error; err != nil {
		return nil, err
	}
	return &api2go.Response{
		Res:  obj,
		Code: http.StatusOK,
	}, nil
}
