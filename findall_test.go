package resource

import (
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
	if err := verifyCount(a, 2); err != nil {
		t.Fatal(err)
	}
}
