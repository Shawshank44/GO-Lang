package utils

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

func GenerateInsertQuery(tableName string, model interface{}) string {
	modelType := reflect.TypeOf(model)
	var columns, placeholders string
	for i := 0; i < modelType.NumField(); i++ {
		dbTag := modelType.Field(i).Tag.Get("db")
		fmt.Println("dbTag", dbTag)
		dbTag = strings.TrimSuffix(dbTag, ",omitempty")
		if dbTag != "" && dbTag != "id" { //skip the id field if its auto increment
			if columns != "" {
				columns += ", "
				placeholders += ", "
			}
			columns += dbTag
			placeholders += "?"
		}
	}
	// fmt.Printf("INSERT INTO teachers (%s) VALUES (%s) \n", columns, placeholders)
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, columns, placeholders)
}

func GetStructValues(model interface{}) []interface{} {
	modelValue := reflect.ValueOf(model)
	modelType := modelValue.Type()
	values := []interface{}{}
	for i := 0; i < modelType.NumField(); i++ {
		dbTag := modelType.Field(i).Tag.Get("db")
		if dbTag != "" && dbTag != "id,omitempty" {
			values = append(values, modelValue.Field(i).Interface())
		}
	}
	fmt.Println("Values from GetStructValues function", values)
	return values
}

func IsValidSortOrder(order string) bool {
	return order == "asc" || order == "desc"
}

func IsValidSortField(field string) bool {
	validFields := map[string]bool{
		"first_name": true,
		"last_name":  true,
		"email":      true,
		"class":      true,
		"subject":    true,
	}
	return validFields[field]
}

func AddSorting(r *http.Request, query string) string {
	sortParams := r.URL.Query()["sortby"] // converts in string slice

	if len(sortParams) > 0 {
		query += " ORDER BY"
		for i, param := range sortParams {
			parts := strings.Split(param, ":")
			if len(parts) != 2 {
				continue
			}
			fields, order := parts[0], parts[1]
			if !IsValidSortField(fields) || !IsValidSortOrder(order) {
				continue
			}
			if i > 0 {
				query += ","
			}
			query += " " + fields + " " + order
		}
	}
	return query
}

func AddFilters(r *http.Request, query string, args []interface{}) (string, []interface{}) {
	params := map[string]string{
		"first_name": "first_name",
		"last_name":  "last_name",
		"email":      "email",
		"class":      "class",
		"subject":    "subject",
	}

	for params, dbfield := range params {
		value := r.URL.Query().Get(params)
		if value != "" {
			query += " AND " + dbfield + " = ?"
			args = append(args, value)
		}
	}
	return query, args
}
