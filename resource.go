package resource

import (
	"github.com/jinzhu/gorm"
)

// Resource implements the interfaces necessary to use a GORM model with the
// api2go package.
type Resource struct {
	DB   *gorm.DB
	Type interface{}
}
