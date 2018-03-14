package resource

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/manyminds/api2go"
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
	_, err = sendRequest(
		a,
		http.MethodGet,
		fmt.Sprintf("/articles/%d", article.ID),
		nil,
		http.StatusOK,
	)
	if err != nil {
		t.Fatal(err)
	}
}
