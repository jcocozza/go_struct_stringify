package gostructstringify

import (
	"fmt"
	"math"
	"reflect"
)

// Return a string representation of a struct instance code.
// Works recursively for all sub structs as well.
//
// i.e. "&myStruct{A: a, B: 1}"
func StructStringify(instance any) string {
	code := ""
	instanceValue := reflect.ValueOf(instance)
	if instanceValue.Kind() == reflect.Ptr {
		// If the input is a pointer, dereference it
		instanceValue = instanceValue.Elem()
		code += "&"
	}

	instanceType := instanceValue.Type()

	if instanceType.Kind() == reflect.Slice {
		code += instanceType.String() +"{"
		for i := 0; i < instanceValue.Len(); i++ {
			if i == instanceValue.Len() - 1 {
				code += StructStringify(instanceValue.Index(i).Interface())
			} else {
				code += StructStringify(instanceValue.Index(i).Interface()) + ", "
			}
		}
		code += "}"
		return code
	}

	/*
	TODO: It feels like there should be a way to generalize these basic type cases
	e.g. type MYCustType string or type MyCustType float64
	*/
	if instanceType.Kind() == reflect.String && reflect.TypeOf(instance) != reflect.TypeOf("string") {
		code += instanceType.String() + "(\"" + fmt.Sprint(instance) + "\")"
		return code
	}
	if instanceType.Kind() == reflect.Float64 && reflect.TypeOf(instance) != reflect.TypeOf(1.00) {
		code += instanceType.String() +  "(" + fmt.Sprint(instance) + ")"
		return code
	}
	// when a float ins a nan
	if instanceType.Kind() == reflect.Float64 &&  math.IsNaN(instanceValue.Float()) {
		code += "math.NaN()"
		return code
	}

	code += instanceType.String() + "{"

	for i := 0; i < instanceType.NumField(); i++ {
		field := instanceType.Field(i)
		fieldValue := instanceValue.Field(i).Interface()

		switch fieldValue := fieldValue.(type) {
		case int, float64, bool: // handle basic types
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