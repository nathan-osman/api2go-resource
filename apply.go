package resource

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/manyminds/api2go"
)

// ErrInvalidParameter indicates that an invalid parameter was supplied to a
// request.
var ErrInvalidParameter = errors.New("invalid parameter")

// apply takes the query parameters from a request and applies them to an SQL
// query.
func (r *Resource) apply(req api2go.Request) (*gorm.DB, error) {
	c := r.DB
loop:
	for k, v := range req.QueryParams {
		for _, f := range r.Fields {
			if f == k {
				c = c.Where(fmt.Sprintf("%s = ?", k), v[0])
				continue loop
			}
		}
		return nil, ErrInvalidParameter
	}
	return c, nil
}
