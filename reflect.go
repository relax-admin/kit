package kit

import (
	"reflect"
)

func FieldByName(o interface{}, name string) interface{} {
	v := reflect.ValueOf(o)
	return v.FieldByName(name)
}
