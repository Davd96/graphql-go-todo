package services

import (
	"fmt"
	"reflect"
)

func makeUpdateQuery(obj interface{}) string {

	rv := reflect.ValueOf(obj)
	opt := ""

	for i := 0; i < rv.NumField(); i++ {

		if !rv.Field(i).IsZero() {
			field := rv.Field(i)

			if field.Kind() == reflect.Ptr {
				field = field.Elem()
			}

			if i != 0 && opt != "" {
				opt += ", "
			}

			switch field.Kind() {
			case reflect.Int32:
				opt += fmt.Sprintf("%s= %v", rv.Type().Field(i).Tag.Get("db"), field)
			case reflect.String:
				opt += fmt.Sprintf("%s= '%s'", rv.Type().Field(i).Tag.Get("db"), field)
			default:
				fmt.Println(field.Kind())
			}
		}
	}

	return opt
}
