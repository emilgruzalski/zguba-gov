package odata

import (
	"fmt"
	"regexp"
	"strings"
)

var allowedFilterFields = map[string]string{
	"item_status":       "item_status",
	"item_category":     "item_category",
	"municipality_name": "municipality_name",
	"item_name":         "item_name",
	"item_description":  "item_description",
}

var allowedOrderFields = map[string]string{
	"created_at": "created_at",
	"item_name":  "item_name",
	"item_date":  "item_date",
}

type FilterClause struct {
	Where string
	Args  []any
}

func ParseFilter(filter string) (*FilterClause, error) {
	if filter == "" {
		return &FilterClause{}, nil
	}

	eqRe := regexp.MustCompile(`(\w+)\s+eq\s+'([^']*)'`)
	containsRe := regexp.MustCompile(`contains\((\w+),\s*'([^']*)'\)`)
	startsWithRe := regexp.MustCompile(`startswith\((\w+),\s*'([^']*)'\)`)

	var clauses []string
	var args []any

	for _, match := range eqRe.FindAllStringSubmatch(filter, -1) {
		col, ok := allowedFilterFields[match[1]]
		if !ok {
			return nil, fmt.Errorf("unsupported filter field: %s", match[1])
		}
		clauses = append(clauses, col+" = ?")
		args = append(args, match[2])
	}

	for _, match := range containsRe.FindAllStringSubmatch(filter, -1) {
		col, ok := allowedFilterFields[match[1]]
		if !ok {
			return nil, fmt.Errorf("unsupported filter field: %s", match[1])
		}
		clauses = append(clauses, col+" LIKE ?")
		args = append(args, "%"+match[2]+"%")
	}

	for _, match := range startsWithRe.FindAllStringSubmatch(filter, -1) {
		col, ok := allowedFilterFields[match[1]]
		if !ok {
			return nil, fmt.Errorf("unsupported filter field: %s", match[1])
		}
		clauses = append(clauses, col+" LIKE ?")
		args = append(args, match[2]+"%")
	}

	return &FilterClause{
		Where: strings.Join(clauses, " AND "),
		Args:  args,
	}, nil
}

func ParseOrderBy(orderby string) string {
	if orderby == "" {
		return "created_at DESC"
	}

	parts := strings.SplitN(strings.TrimSpace(orderby), " ", 2)
	col, ok := allowedOrderFields[parts[0]]
	if !ok {
		return "created_at DESC"
	}

	dir := "ASC"
	if len(parts) > 1 && strings.EqualFold(parts[1], "desc") {
		dir = "DESC"
	}

	return col + " " + dir
}
