package resource

import (
	"net/http"
	"reflect"

	"github.com/manyminds/api2go"
)

// FindAll attempts to retrieve all instances of a model from the database.
func (r *Resource) FindAll(req api2go.Request) (api2go.Responder, error) {
	p := &Params{
		Action:  BeforeFindAll,
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
		objType   = reflect.TypeOf(r.Type)
		sliceType = reflect.SliceOf(objType)
		sliceVal  = reflect.New(sliceType)
	)
	if p.DB = p.DB.Find(sliceVal.Interface()); p.DB.Error != nil {
		return nil, p.DB.Error
	}
	p.Action = AfterFindAll
	if err := r.runHooks(p); err != nil {
		return nil, err
	}
	return &api2go.Response{
		Res:  reflect.Indirect(sliceVal).Interface(),
		Code: http.StatusOK,
	}, nil
}
