package resource

import (
	"net/http"
	"reflect"

	"github.com/manyminds/api2go"
)

// FindOne attempts to retrieve a single model instance from the database.
func (r *Resource) FindOne(ID string, req api2go.Request) (api2go.Responder, error) {
	c, err := r.apply(req)
	if err != nil {
		return nil, err
	}
	var (
		objType = reflect.TypeOf(r.Type).Elem()
		objVal  = reflect.New(objType)
	)
	if err := c.First(objVal.Interface(), ID).Error; err != nil {
		return nil, err
	}
	return &api2go.Response{
		Res:  objVal.Interface(),
		Code: http.StatusOK,
	}, nil
}
