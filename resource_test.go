package resource

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/manyminds/api2go"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Article struct {
	ID      int64  `json:"id,omitempty"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (a *Article) GetName() string {
	return "articles"
}

func (a *Article) GetID() string {
	return strconv.FormatInt(a.ID, 10)
}

func (a *Article) SetID(id string) error {
	a.ID, _ = strconv.ParseInt(id, 10, 64)
	return nil
}

var (
	Article1 = &Article{
		Title:   "Title 1",
		Content: "Content 1",
	}
	Article2 = &Article{
		Title:   "Title 2",
		Content: "Content 2",
	}
)

func initDatabase() (*gorm.DB, *api2go.API, error) {
	c, err := gorm.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		return nil, nil, err
	}
	if err := c.AutoMigrate(Article1).Error; err != nil {
		c.Close()
		return nil, nil, err
	}
	return c, api2go.NewAPI(""), nil
}
