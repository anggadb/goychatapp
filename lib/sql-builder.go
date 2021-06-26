package lib

import (
	"reflect"
	"strconv"
)

func DynamicFilters(ix interface{}, softDeleted bool) (string, error) {
	where := ""
	if softDeleted {
		where = "deleted_at IS NULL AND "
	}
	values := reflect.ValueOf(ix)
	args := []interface{}{}
	for i := 0; i < values.NumField(); i++ {
		f := values.Field(i)
		if f.IsZero() || !f.IsValid() || f.IsNil() {
			continue
		}
		v := f.Elem().Interface()
		switch ks := v.(type) {
		case int:
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
		col := values.Type().Field(i).Tag.Get("json")
		where += col + "=$" + strconv.Itoa(len(args)) + " AND "
	}
	return where, nil
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	m := map[string]interface{}{"UserID": 1234, "Age": 18}
// var values []interface{}
// var where []string
// for k, v := range m {
//     values = append(values, v)
//     where = append(where, fmt.Sprintf("%s = ?", k))
// }
// fmt.Println(where)
// }
