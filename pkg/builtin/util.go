package builtin

import "reflect"

func GetClassName(obj GDClass) string {
	return reflect.TypeOf(obj).Elem().Name()
}
