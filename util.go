package resource

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/manyminds/api2go"
)

// ErrInvalidParameter indicates that an invalid parameter was supplied to a
// request.
var ErrInvalidParameter = errors.New("invalid parameter")

// apply takes the query parameters from a request and applies them to an SQL
// query.
func (r *Resource) apply(p *Params) error {
loop:
	for k, v := range p.Request.QueryParams {
		for _, f := range r.Fields {
			if f == k {
				p.DB = p.DB.Where(fmt.Sprintf("%s = ?", k), v[0])
				continue loop
			}
		}
		return ErrInvalidParameter
	}
	return nil
}

// translateError takes a database query and ensures an appropriate HTTPError
// is returned when the query fails.
func translateError(db *gorm.DB) error {
	if db.Error != nil {
		if db.RecordNotFound() {
			return api2go.NewHTTPError(
				nil,
				http.StatusText(http.StatusNotFound),
				http.StatusNotFound,
			)
		}
		return db.Error
	}
	return nil
}
