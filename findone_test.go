package resource

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFindOne(t *testing.T) {
	c, a, err := initDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	article, err := createArticle(c)
	if err != nil {
		t.Fatal(err)
	}
	a.AddResource(&Article{}, &Resource{
		DB:   c,
		Type: &Article{},
	})
	r, err := testRequest(
		a,
		http.MethodGet,
		fmt.Sprintf("/articles/%d", article.ID),
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
	if d.Data.DataObject == nil {
		t.Fatal("object is nil")
	}
}
