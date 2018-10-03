package resource

import (
	"net/http"

	"github.com/manyminds/api2go"
)

// Delete attempts to delete the specified model instance from the database.
func (r *Resource) Delete(id string, req api2go.Request) (api2go.Responder, error) {
	p := &Params{
		Action:  Delete,
		Request: req,
		DB:      r.DB,
	}
	if err := r.runHooks(p); err != nil {
		return nil, err
	}
	if err := p.DB.Where("ID = ?", id).Delete(r.Type).Error; err != nil {
		return nil, err
	}
	return &api2go.Response{
		Code: http.StatusOK,
	}, nil
}
