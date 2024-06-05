package generate_query

import (
	"fmt"
	"reflect"
	"strings"

	repositoryQuery "github.com/ShevelevEvgeniy/app/internal/repository/repository_query"
)

func GenerateInsertQuery[T interface{}](model T, tableName string, onConflictUpdate bool) (string, []interface{}, error) {
	if onConflictUpdate {
		return generateUpdateQuery(model, tableName)
	}

	return generateSaveQuery(model, tableName)
}

func generateColumnsAndValues[T interface{}](model T) ([]string, []interface{}, []string) {
	var columns []string
	var values []interface{}
	var placeholders []string

	val := reflect.ValueOf(model)

	for j := 0; j < val.NumField(); j++ {
		jsonTag := val.Type().Field(j).Tag.Get("json")

		if jsonTag == Id || jsonTag == updatedAt || jsonTag == createdAt {
			continue
		}

		columns = append(columns, jsonTag)
		values = append(values, val.Field(j).Interface())

		if val.Type().Field(j).Tag.Get("description") == Wkb {
			placeholders = append(placeholders, fmt.Sprintf("%s($%d)", Wkb, len(values)))
			continue
		}

		placeholders = append(placeholders, fmt.Sprintf("$%d", len(values)))
	}

	return columns, values, placeholders
}

func generateSaveQuery[T interface{}](model T, tableName string) (string, []interface{}, error) {
	columns, values, placeholders := generateColumnsAndValues(model)

	query := fmt.Sprintf(repositoryQuery.SaveOrUpdate,
		tableName,
		fmt.Sprintf("(%s)", strings.Join(columns, ", ")),
		fmt.Sprintf("(%s)", strings.Join(placeholders, ", ")),
	)

	return query, values, nil
}

func generateUpdateQuery[T interface{}](model T, tableName string) (string, []interface{}, error) {
	columns, values, placeholders := generateColumnsAndValues(model)

	var setValues []string
	for _, col := range columns {
		setValues = append(setValues, fmt.Sprintf("%s = EXCLUDED.%s", col, col))
	}

	var uniqueField string
	for j := 0; j < reflect.TypeOf(model).NumField(); j++ {
		if reflect.TypeOf(model).Field(j).Tag.Get("unique") == unique {
			uniqueField = reflect.TypeOf(model).Field(j).Tag.Get("json")
			break
		}
	}

	query := fmt.Sprintf(repositoryQuery.SaveOrUpdate,
		tableName,
		fmt.Sprintf("(%s)", strings.Join(columns, ", ")),
		fmt.Sprintf("(%s)", strings.Join(placeholders, ", ")),
		uniqueField,
		strings.Join(setValues, ", "),
	)

	return query, values, nil
}
