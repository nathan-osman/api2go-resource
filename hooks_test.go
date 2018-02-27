package resource

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/manyminds/api2go"
)

func TestGetHook(t *testing.T) {
	c, a, err := initDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	for i := 0; i < 3; i++ {
		_, err := createArticle(c)
		if err != nil {
			t.Fatal(err)
		}
	}
	a.AddResource(&Article{}, &Resource{
		DB:   c,
		Type: &Article{},
		GetHooks: []GetHook{
			func(c *gorm.DB, req api2go.Request) *gorm.DB {
				return c.Limit("2")
			},
		},
	})
	if err := verifyCount(a, 2); err != nil {
		t.Fatal(err)
	}
}
