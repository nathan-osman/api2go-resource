package resource

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/manyminds/api2go/jsonapi"
)

func TestHooks(t *testing.T) {
	c, a, err := initDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	article, err := createArticle(c)
	if err != nil {
		t.Fatal(err)
	}
	hookInvoked := false
	a.AddResource(&Article{}, &Resource{
		DB:   c,
		Type: &Article{},
		Hooks: []Hook{
			func(*Params) error {
				hookInvoked = true
				return nil
			},
		},
	})
	b, err := jsonapi.Marshal(article)
	if err != nil {
		t.Fatal(err)
	}
	_, err = sendRequest(
		a,
		http.MethodPatch,
		fmt.Sprintf("/articles/%d", article.ID),
		bytes.NewReader(b),
		http.StatusOK,
	)
	if err != nil {
		t.Fatal(err)
	}
	if !hookInvoked {
		t.Fatal("hook was not invoked")
	}
}
