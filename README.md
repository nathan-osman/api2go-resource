## api2go-resource

[![GoDoc](https://godoc.org/github.com/nathan-osman/api2go-resource?status.svg)](https://godoc.org/github.com/nathan-osman/api2go-resource)
[![MIT License](http://img.shields.io/badge/license-MIT-9370d8.svg?style=flat)](http://opensource.org/licenses/MIT)

This package serves as a bridge between [GORM](https://github.com/jinzhu/gorm) and [api2go](https://github.com/manyminds/api2go), reducing the amount of boilerplate code needed for implementing CRUD actions for GORM models.

### Usage

Let's suppose you have the following model definition:

    type Article struct {
        ID      int64
        Title   string
        Content string
    }

In order to use the model with api2go, three important methods must be implemented:

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

The next step is to create a `Resource` instance for the model:

    import "github.com/nathan-osman/api2go-resource"

    // db is an instance of *gorm.DB

    var articleResource = &resource.Resource{
        DB:   db,
        Type: &Article{},
    }

This resource can now be registered with api2go:

    api := api2go.NewAPI("api")
    api.AddResource(&Article{}, articleResource)
