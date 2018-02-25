package resource

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/manyminds/api2go/jsonapi"
)

func TestUpdate(t *testing.T) {
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
	b, err := jsonapi.Marshal(article)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := testRequest(
		a,
		http.MethodPatch,
		fmt.Sprintf("/articles/%d", article.ID),
		bytes.NewReader(b),
		http.StatusOK,
	); err != nil {
		t.Fatal(err)
	}
}
