package menu

import (
	"fmt"
	"strings"
)

func BuildMenuCacheKey(filter MenuFilter) string {

	var builder strings.Builder

	builder.WriteString("menus")

	if filter.Search != "" {
		builder.WriteString(":search=")
		builder.WriteString(strings.ToLower(filter.Search))
	}

	if filter.Available != nil {
		builder.WriteString(fmt.Sprintf(
			":available=%t",
			*filter.Available,
		))
	}

	builder.WriteString(fmt.Sprintf(
		":page=%d",
		filter.Page,
	))

	builder.WriteString(fmt.Sprintf(
		":limit=%d",
		filter.Limit,
	))

	builder.WriteString(":sort=")
	builder.WriteString(filter.SortBy)

	builder.WriteString(":order=")
	builder.WriteString(filter.Order)

	return builder.String()

}
