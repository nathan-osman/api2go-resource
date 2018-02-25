package resource

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDelete(t *testing.T) {
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
	if _, err := testRequest(
		a,
		http.MethodDelete,
		fmt.Sprintf("/articles/%d", article.ID),
		nil,
		http.StatusOK,
	); err != nil {
		t.Fatal(err)
	}
}
