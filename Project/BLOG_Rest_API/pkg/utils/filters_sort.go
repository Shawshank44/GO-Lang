package utils

import (
	"net/http"
	"strings"
)

func IsValidSortOrder(order string) bool {
	return order == "asc" || order == "desc"
}

func IsValidSortField(field string) bool {
	validFields := map[string]bool{
		"username": true,
		"email":    true,
	}
	return validFields[field]
}

func AddSorting(r *http.Request, query string) string {
	sortParams := r.URL.Query()["sortby"]
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
		"username": "username",
		"email":    "email",
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
