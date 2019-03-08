package migration

import (
	"reflect"
	"strings"
)

const structTag = "crate"

func mapGoToCrateType(value string) (typeName string) {
	switch value {
	case "bool":
		typeName = "boolean"
	case "string":
		typeName = "string"
	case "Time":
		typeName = "timestamp"
	case "int":
		fallthrough
	case "int32":
		fallthrough
	case "int64":
		typeName = "integer"
	case "float32":
		typeName = "float"
	case "float64":
		typeName = "double"
	}
	return
}

func GetColumnNames(v interface{}) []string {
	t := reflect.TypeOf(v)
	columns := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldName := field.Name

		tag := field.Tag.Get("crate")

		values := strings.Split(tag, ";")
		for _, opt := range values {
			keyValue := strings.Split(opt, ":")
			if len(keyValue) != 2 {
				continue
			}

			if keyValue[0] == "column" {
				fieldName = keyValue[1]
			}
		}

		columns[i] = fieldName
	}

	return columns
}

func GetSchemaForCrate(v interface{}) string {
	t := reflect.TypeOf(v)

	tableName := TableName(v)
	createStatement := "CREATE TABLE IF NOT EXISTS " + tableName + " ( "

	fields := make([]string, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		typeName := mapGoToCrateType(field.Type.Name())

		fieldName := field.Name

		tag := field.Tag.Get("crate")

		values := strings.Split(tag, ";")
		for _, opt := range values {
			keyValue := strings.Split(opt, ":")
			if len(keyValue) != 2 {
				continue
			}

			switch keyValue[0] {
			case "column":
				fieldName = keyValue[1]
			case "type":
				typeName = keyValue[1]
			}
		}

		fields[i] = fieldName + " " + typeName
	}

	createStatement += strings.Join(fields, " , ")
	createStatement += " );"
	return createStatement
}
