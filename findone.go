package resource

import (
	"net/http"
	"reflect"

	"github.com/manyminds/api2go"
)

// FindOne attempts to retrieve a single model instance from the database.
func (r *Resource) FindOne(ID string, req api2go.Request) (api2go.Responder, error) {
	p := &Params{
		Action:  BeforeFindOne,
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
	if p.DB = p.DB.First(objVal.Interface(), ID); p.DB.Error != nil {
		return nil, p.DB.Error
	}
	p.Action = AfterFindOne
	if err := r.runHooks(p); err != nil {
		return nil, err
	}
	return &api2go.Response{
		Res:  objVal.Interface(),
		Code: http.StatusOK,
	}, nil
}
