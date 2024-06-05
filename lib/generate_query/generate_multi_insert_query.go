package generate_query

import (
	"fmt"
	"reflect"
	"strings"

	repositoryQuery "github.com/ShevelevEvgeniy/app/internal/repository/repository_query"
)

func GenerateMultiInsertQuery[T interface{}](models []T, tableName string, onConflictUpdate bool) (string, []interface{}, error) {
	if onConflictUpdate {
		return generateMultiQuery(models, tableName, true)
	}
	return generateMultiQuery(models, tableName, false)
}

func generateMultiQuery[T interface{}](models []T, tableName string, isUpdate bool) (string, []interface{}, error) {
	var values []interface{}
	var columns []string
	var placeholders string

	for i, model := range models {
		var groupPlaceholders []string
		var valuesGroup []interface{}

		if i == 0 {
			columns, valuesGroup, groupPlaceholders = generateColumnsAndValuesForMultiInsert(model, len(values))
		} else {
			_, valuesGroup, groupPlaceholders = generateColumnsAndValuesForMultiInsert(model, len(values))
		}

		values = append(values, valuesGroup...)

		placeholders += fmt.Sprintf("(%s), ", strings.Join(groupPlaceholders, ", "))
	}

	query := ""
	if isUpdate {
		query = generateMultiUpdateQuery(models[0], tableName, columns, placeholders)
	} else {
		query = generateMultiInsertQuery(tableName, columns, placeholders)
	}

	return query, values, nil
}

func generateMultiInsertQuery(tableName string, columns []string, placeholders string) string {
	return fmt.Sprintf(repositoryQuery.Save,
		tableName,
		fmt.Sprintf("(%s)", strings.Join(columns, ", ")),
		strings.TrimRight(placeholders, ", ")+";",
	)
}

func generateMultiUpdateQuery(model interface{}, tableName string, columns []string, placeholders string) string {
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

	return fmt.Sprintf(repositoryQuery.SaveOrUpdate,
		tableName,
		fmt.Sprintf("(%s)", strings.Join(columns, ", ")),
		strings.TrimRight(placeholders, ", "),
		uniqueField,
		strings.Join(setValues, ", ")+";",
	)
}

func generateColumnsAndValuesForMultiInsert[T interface{}](model T, valueIndex int) ([]string, []interface{}, []string) {
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
			placeholders = append(placeholders, fmt.Sprintf("%s($%d)", Wkb, valueIndex+len(values)))
			continue
		}

		placeholders = append(placeholders, fmt.Sprintf("$%d", valueIndex+len(values)))
	}

	return columns, values, placeholders
}
