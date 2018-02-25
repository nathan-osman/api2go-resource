package resource

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/manyminds/api2go/jsonapi"
)

func TestCreate(t *testing.T) {
	c, a, err := initDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	a.AddResource(&Article{}, &Resource{
		DB:   c,
		Type: &Article{},
	})
	b, err := jsonapi.Marshal(&Article{})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := testRequest(
		a,
		http.MethodPost,
		"/articles",
		bytes.NewReader(b),
		http.StatusCreated,
	); err != nil {
		t.Fatal(err)
	}
}
