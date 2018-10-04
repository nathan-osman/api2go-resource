package resource

import (
	"net/http"

	"github.com/manyminds/api2go"
)

// Delete attempts to delete the specified model instance from the database.
func (r *Resource) Delete(id string, req api2go.Request) (api2go.Responder, error) {
	p := &Params{
		Action:  BeforeDelete,
		Request: req,
		DB:      r.DB,
	}
	if err := r.runHooks(p); err != nil {
		return nil, err
	}
	if p.DB = p.DB.Where("ID = ?", id).Delete(r.Type); p.DB.Error != nil {
		return nil, p.DB.Error
	}
	p.Action = AfterDelete
	if err := r.runHooks(p); err != nil {
		return nil, err
	}
	return &api2go.Response{
		Code: http.StatusOK,
	}, nil
}
