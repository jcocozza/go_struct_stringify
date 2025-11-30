package gostructstringify

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"
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
		code += instanceType.String() + "{"
		for i := 0; i < instanceValue.Len(); i++ {
			if i == instanceValue.Len()-1 {
				code += StructStringify(instanceValue.Index(i).Interface())
			} else {
				code += StructStringify(instanceValue.Index(i).Interface()) + ", "
			}
		}
		code += "}"
		return code
	}

	switch instanceType.Kind() {
	case reflect.String:
		return fmt.Sprintf("%q", instanceValue.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", instanceValue.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", instanceValue.Uint())
	case reflect.Float32, reflect.Float64:
		if math.IsNaN(instanceValue.Float()) {
			return "math.NaN()"
		}
		return fmt.Sprintf("%v", instanceValue.Float())
	case reflect.Bool:
		return fmt.Sprintf("%v", instanceValue.Bool())
	}

	code += instanceType.String() + "{"

	for i := 0; i < instanceType.NumField(); i++ {
		field := instanceType.Field(i)

		if field.PkgPath != "" { // unexported
			ptr := unsafe.Pointer(instanceValue.UnsafeAddr())
			fieldPtr := unsafe.Pointer(uintptr(ptr) + field.Offset)
			fieldVal := reflect.NewAt(field.Type, fieldPtr).Elem()
			code += fmt.Sprintf("%s: %v, ", field.Name, StructStringify(fieldVal.Interface()))
			continue
		}
		fieldValue := instanceValue.Field(i).Interface()
		code += fmt.Sprintf("%s: %s, ", field.Name, StructStringify(fieldValue))
	}
	code = code[:len(code)-2] + "}"
	return code
}
