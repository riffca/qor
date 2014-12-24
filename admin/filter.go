package admin

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/qor/qor"
)

type Filter struct {
	Name       string
	Operations []string
	Handler    func(name string, value string, scope *gorm.DB, context *qor.Context) *gorm.DB
}

var DefaultHandler = func(name string, value string, scope *gorm.DB, context *qor.Context) *gorm.DB {
	lastIndex := strings.LastIndex(name, "_")
	operation := name[lastIndex+1 : len(name)]
	column := name[0:lastIndex]

	switch operation {
	case "cont":
		return scope.Where(fmt.Sprintf("%v ILIKE ?", scope.NewScope(nil).Quote(column)), "%"+value+"%")
	case "eq":
		return scope.Where(fmt.Sprintf("%v = ?", scope.NewScope(nil).Quote(column)), value)
	case "gt":
		return scope.Where(fmt.Sprintf("%v > ?", scope.NewScope(nil).Quote(column)), value)
	case "gteq":
		return scope.Where(fmt.Sprintf("%v >= ?", scope.NewScope(nil).Quote(column)), value)
	case "lt":
		return scope.Where(fmt.Sprintf("%v < ?", scope.NewScope(nil).Quote(column)), value)
	case "lteq":
		return scope.Where(fmt.Sprintf("%v <= ?", scope.NewScope(nil).Quote(column)), value)
	}
	return scope
}
