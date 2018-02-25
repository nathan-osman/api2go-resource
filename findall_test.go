package resource

import (
	"net/http"
	"testing"
)

func TestFindAll(t *testing.T) {
	c, a, err := initDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	for i := 0; i < 2; i++ {
		_, err := createArticle(c)
		if err != nil {
			t.Fatal(err)
		}
	}
	a.AddResource(&Article{}, &Resource{
		DB:   c,
		Type: &Article{},
	})
	r, err := testRequest(
		a,
		http.MethodGet,
		"/articles",
		nil,
		http.StatusOK,
	)
	if err != nil {
		t.Fatal(err)
	}
	d, err := unmarshalResponse(r)
	if err != nil {
		t.Fatal(err)
	}
	if len(d.Data.DataArray) != 2 {
		t.Fatalf("%d != %d", len(d.Data.DataArray), 2)
	}
}
