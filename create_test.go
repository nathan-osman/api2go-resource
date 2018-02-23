package resource

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manyminds/api2go/jsonapi"
)

func TestCreate(t *testing.T) {
	c, a, err := initDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	a.AddResource(Article1, &Resource{
		DB:   c,
		Type: Article1,
	})
	b, err := jsonapi.Marshal(Article1)
	if err != nil {
		t.Fatal(err)
	}
	var (
		w = httptest.NewRecorder()
		r = httptest.NewRequest(
			http.MethodPost,
			"/articles",
			bytes.NewReader(b),
		)
	)
	a.Handler().ServeHTTP(w, r)
	if w.Code != http.StatusCreated {
		t.Fatalf("%d != %d", w.Code, http.StatusCreated)
	}
}
