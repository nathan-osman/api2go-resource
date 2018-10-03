package resource

import (
	"net/http"
	"reflect"

	"github.com/manyminds/api2go"
)

// FindOne attempts to retrieve a single model instance from the database.
func (r *Resource) FindOne(ID string, req api2go.Request) (api2go.Responder, error) {
	p := &Params{
		Action:  FindOne,
		Request: req,
		DB:      r.DB,
	}
	if err := r.apply(p); err != nil {
		return nil, err
	}
	if err := r.runHooks(p); err != nil {
		return nil, err
	}
	var (
		objType = reflect.TypeOf(r.Type).Elem()
		objVal  = reflect.New(objType)
	)
	if db := p.DB.First(objVal.Interface(), ID); db.Error != nil {
		if db.RecordNotFound() {
			return nil, api2go.NewHTTPError(
				nil,
				http.StatusText(http.StatusNotFound),
				http.StatusNotFound,
			)
		}
		return nil, db.Error
	}
	return &api2go.Response{
		Res:  objVal.Interface(),
		Code: http.StatusOK,
	}, nil
}
