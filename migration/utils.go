package migration

import (
	"reflect"
	"strings"

	"github.com/gedex/inflector"
)

func TableName(v interface{}) string {
	t := reflect.TypeOf(v)
	tableName := strings.ToLower(t.Name())
	return inflector.Pluralize(tableName)
}
