package resource

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/manyminds/api2go"
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
	var (
		globalHookInvoked = false
		getHookInvoked    = false
		setHookInvoked    = false
	)
	a.AddResource(&Article{}, &Resource{
		DB:   c,
		Type: &Article{},
		GlobalHooks: []GlobalHook{
			func(Action, api2go.Request) error {
				globalHookInvoked = true
				return nil
			},
		},
		GetHooks: []GetHook{
			func(c *gorm.DB, req api2go.Request) *gorm.DB {
				getHookInvoked = true
				return c
			},
		},
		SetHooks: []SetHook{
			func(interface{}, api2go.Request) {
				setHookInvoked = true
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
	if !globalHookInvoked {
		t.Fatal("global hook was not invoked")
	}
	if !getHookInvoked {
		t.Fatal("get hook was not invoked")
	}
	if !setHookInvoked {
		t.Fatal("set hook was not invoked")
	}
}
