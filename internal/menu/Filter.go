package menu

import (
	"fmt"
	"strings"
)

type MenuFilter struct {
	Search string

	Available *bool

	Page int

	Limit int

	SortBy string

	Order string
}

func (f *MenuFilter) Normalize() {

	if f.Page <= 0 {
		f.Page = 1
	}

	if f.Limit <= 0 {
		f.Limit = 10
	}

	if f.Limit > 100 {
		f.Limit = 100
	}

	allowedSort := map[string]bool{
		"name":       true,
		"price":      true,
		"created_at": true,
	}

	if !allowedSort[f.SortBy] {
		f.SortBy = "name"
	}

	f.Order = strings.ToLower(f.Order)

	if f.Order != "asc" && f.Order != "desc" {
		f.Order = "asc"
	}
}

func (f MenuFilter) OrderClause() string {
	return fmt.Sprintf("%s %s", f.SortBy, f.Order)
}
