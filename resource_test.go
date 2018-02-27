package resource

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/manyminds/api2go"
	"github.com/manyminds/api2go/jsonapi"

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

func initDatabase() (*gorm.DB, *api2go.API, error) {
	c, err := gorm.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		return nil, nil, err
	}
	if err := c.AutoMigrate(&Article{}).Error; err != nil {
		c.Close()
		return nil, nil, err
	}
	return c, api2go.NewAPI(""), nil
}

func createArticle(c *gorm.DB) (*Article, error) {
	a := &Article{
		Title:   "Title",
		Content: "Content",
	}
	if err := c.Create(a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

func testRequest(a *api2go.API, method, target string, body io.Reader, code int) (*http.Response, error) {
	var (
		w = httptest.NewRecorder()
		r = httptest.NewRequest(method, target, body)
	)
	a.Handler().ServeHTTP(w, r)
	if w.Code != code {
		return nil, fmt.Errorf("%d != %d", w.Code, code)
	}
	return w.Result(), nil
}

func unmarshalResponse(r *http.Response) (*jsonapi.Document, error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	doc := &jsonapi.Document{}
	if err := json.Unmarshal(b, doc); err != nil {
		return nil, err
	}
	return doc, nil
}

func verifyCount(a *api2go.API, count int) error {
	r, err := testRequest(
		a,
		http.MethodGet,
		"/articles",
		nil,
		http.StatusOK,
	)
	if err != nil {
		return err
	}
	d, err := unmarshalResponse(r)
	if err != nil {
		return err
	}
	if len(d.Data.DataArray) != 2 {
		return fmt.Errorf("%d != %d", len(d.Data.DataArray), 2)
	}
	return nil
}
