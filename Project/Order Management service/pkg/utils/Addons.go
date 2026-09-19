package utils

import (
	"net/http"
	"strconv"
	"strings"
)

func IsValidSortOrder(order string) bool {
	return order == "asc" || order == "desc"
}

func IsValidSortField(field string) bool {
	validFields := map[string]bool{
		"username":     true,
		"email":        true,
		"id":           true,
		"sku":          true,
		"name":         true,
		"category":     true,
		"brand":        true,
		"manufacturer": true,
		"price":        true,
		"currency":     true,
		"stock":        true,
		"unit":         true,
		"status":       true,
		"created_at":   true,
		"updated_at":   true,
	}
	return validFields[field]
}

func AddSorting(r *http.Request, query string) string {
	sortParams := r.URL.Query()["sortby"]

	var sorting []string

	for _, param := range sortParams {
		parts := strings.Split(param, ":")
		if len(parts) != 2 {
			continue
		}
		field := parts[0]
		order := parts[1]

		if !IsValidSortField(field) || !IsValidSortOrder(order) {
			continue
		}
		sorting = append(sorting, field+" "+order)
	}

	if len(sorting) > 0 {
		query += " ORDER BY " + strings.Join(sorting, ", ")
	}

	return query
}

func AddFilters(r *http.Request, query string, args []interface{}) (string, []interface{}) {
	params := map[string]string{
		"username":     "username",
		"email":        "email",
		"sku":          "sku",
		"name":         "name",
		"category":     "category",
		"brand":        "brand",
		"manufacturer": "manufacturer",
		"currency":     "currency",
		"unit":         "unit",
		"status":       "status",
	}

	for param, dbfield := range params {
		value := r.URL.Query().Get(param)

		if value == "" {
			continue
		}
		query += " AND " + dbfield + " = ?"
		args = append(args, value)
	}
	return query, args
}

func GetPaginationParams(r *http.Request) (int, int) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	if limit > 100 {
		limit = 100
	}

	return page, limit
}

func AddSearch(r *http.Request, query string, args []interface{}) (string, []interface{}) {
	search := strings.TrimSpace(r.URL.Query().Get("q"))

	if search == "" {
		return query, args
	}

	searchPattern := "%" + search + "%"

	query += `
		AND (
			sku LIKE ?
			OR name LIKE ?
			OR description LIKE ?
			OR category LIKE ?
			OR brand LIKE ?
			OR manufacturer LIKE ?
		)
	`

	args = append(
		args,
		searchPattern,
		searchPattern,
		searchPattern,
		searchPattern,
		searchPattern,
		searchPattern,
	)

	return query, args
}
