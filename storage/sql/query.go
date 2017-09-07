package sql

import (
	"fmt"
	"reflect"
	"strings"

	_storage "github.com/servicekit/dddkit/storage"
)

func TranslateQuery(query *_storage.Query) string {
	var filterString string
	var orderString string
	var limitString string
	var filterList []string
	var orderList []string

	for _, filter := range query.Filters {
		valueType := "string"
		if reflect.TypeOf(filter.Value).Kind() != reflect.String {
			valueType = ""
		}

		switch filter.Condition {
		case _storage.Equal:
			if valueType == "string" {
				filterList = append(
					filterList,
					fmt.Sprintf("%s='%s'", filter.Property, filter.Value))
			} else {
				filterList = append(
					filterList,
					fmt.Sprintf("%s=%s", filter.Property, filter.Value))
			}
		case _storage.LessThan:
			filterList = append(
				filterList,
				fmt.Sprintf("%s<%s", filter.Property, filter.Value))
		case _storage.LessThanOrEqual:
			filterList = append(
				filterList,
				fmt.Sprintf("%s<=%s", filter.Property, filter.Value))
		case _storage.GreaterThan:
			filterList = append(
				filterList,
				fmt.Sprintf("%s>%s", filter.Property, filter.Value))
		case _storage.GreaterThanOrEqual:
			filterList = append(
				filterList,
				fmt.Sprintf("%s>=%s", filter.Property, filter.Value))
		}
	}

	if len(filterList) > 0 {
		filterString = fmt.Sprintf("WHERE %s", strings.Join(filterList, " AND "))
	}

	for _, order := range query.Orders {
		switch order.Direction {
		case _storage.Ascending:
			orderList = append(
				orderList,
				fmt.Sprintf("%s ASC", order.Property))
		case _storage.Descending:
			orderList = append(
				orderList,
				fmt.Sprintf("%s DESC", order.Property))
		}
	}

	if len(orderList) > 0 {
		orderString = fmt.Sprintf("ORDER BY %s", strings.Join(orderList, ", "))
	}

	if query.Limit > 0 {
		limitString = fmt.Sprintf("LIMIT %d, %d", query.Limit, query.Offset)
	}

	return fmt.Sprintf("%s %s %s", filterString, orderString, limitString)
}
