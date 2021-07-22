package lib

import (
	"reflect"
	"strconv"
)

func DynamicFilters(ix interface{}, softDeleted bool) (string, []interface{}, error) {
	where := ""
	args := []interface{}{}
	if softDeleted {
		where = "deleted_at IS NULL AND "
	}
	rv := reflect.ValueOf(ix)
	for i := 0; i < rv.NumField(); i++ {
		f := rv.Field(i)
		if !f.IsValid() || f.IsNil() {
			continue
		}
		v := f.Elem().Interface()
		switch ks := v.(type) {
		case int:
			if ks == 0 {
				continue
			}
		case uint:
			if ks == 0 {
				continue
			}
		case string:
			if len(ks) == 0 {
				continue
			}
		default:
			continue
		}
		args = append(args, v)
		col := rv.Type().Field(i).Tag.Get("col")
		where += col + " = $" + strconv.Itoa(len(args)) + " AND "
	}
	if wlen := len(where); wlen > 0 {
		where = "WHERE " + where[:wlen-len(" AND ")]
	}
	return where, args, nil
}
