package gostructstringify

import (
	"fmt"
	"reflect"
)

// Return a string representation of a struct instance code.
//
// Works recursively for all sub structs as well
//
// i.e. "&myStruct{ A: a, B: 1}"
func StructStringify(instance any) string {
	instanceValue := reflect.ValueOf(instance)
	if instanceValue.Kind() == reflect.Ptr {
		// If the input is a pointer, dereference it
		instanceValue = instanceValue.Elem()
	}

	instanceType := instanceValue.Type()

	if instanceType.Kind() == reflect.Slice {
		code := instanceType.String() +"{"
		for i := 0; i < instanceValue.Len(); i++ {
			code += StructStringify(instanceValue.Index(i).Interface()) + ", "
		}
		code = code[:len(code)-2] + "}"
		return code
	}

	code := "&" + instanceType.String() + "{"

	for i := 0; i < instanceType.NumField(); i++ {
		field := instanceType.Field(i)
		fieldValue := instanceValue.Field(i).Interface()

		switch fieldValue := fieldValue.(type) {
		case int, float64: // handle basic types
			code += fmt.Sprintf("%s: %v, ", field.Name, fieldValue)
		case string:
			code += fmt.Sprintf("%s: \"%v\", ", field.Name, fieldValue)
		case fmt.Stringer: // handle types implementing Stringer
			code += fmt.Sprintf("%s: %v, ", field.Name, fieldValue)
		default: // handle nested structs
			code += fmt.Sprintf("%s: %v, ", field.Name, StructStringify(fieldValue))
		}
	}

	code = code[:len(code)-2] + "}"

	return code
}