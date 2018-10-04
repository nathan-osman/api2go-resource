package resource

import (
	"fmt"
	"net/http"

	"github.com/manyminds/api2go"
)

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
		return api2go.NewHTTPError(
			nil,
			http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest,
		)
	}
	return nil
}
