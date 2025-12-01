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
	var addr reflect.Value

	if instanceValue.Kind() == reflect.Ptr {
		if instanceValue.IsNil() {
			return fmt.Sprintf("(*%s)(nil)", instanceValue.Type().Elem())
		}
		addr = instanceValue
		instanceValue = instanceValue.Elem() // dereference
		code += "&"
	} else { // when not a pointer we need to get pointer to it
		addr = reflect.New(instanceValue.Type())
		addr.Elem().Set(instanceValue)
		instanceValue = addr.Elem()
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

	// this is an aberration
	switch instanceType.Kind() {
	case reflect.String:
		if reflect.TypeOf(instance) != reflect.TypeOf("") {
			return fmt.Sprintf("%s(%q)", instanceType.String(), instanceValue.String())
		}
		return fmt.Sprintf("%q", instanceValue.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if reflect.TypeOf(instance) != reflect.TypeOf(int(0)) && reflect.TypeOf(instance) != reflect.TypeOf(int8(0)) && reflect.TypeOf(instance) != reflect.TypeOf(int16(0)) && reflect.TypeOf(instance) != reflect.TypeOf(int32(0)) && reflect.TypeOf(instance) != reflect.TypeOf(int64(0)) {
			return fmt.Sprintf("%s(%d)", instanceType.String(), instanceValue.Int())
		}
		return fmt.Sprintf("%d", instanceValue.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if reflect.TypeOf(instance) != reflect.TypeOf(uint(0)) && reflect.TypeOf(instance) != reflect.TypeOf(uint8(0)) && reflect.TypeOf(instance) != reflect.TypeOf(uint16(0)) && reflect.TypeOf(instance) != reflect.TypeOf(uint32(0)) && reflect.TypeOf(instance) != reflect.TypeOf(uint64(0)) {
			return fmt.Sprintf("%s(%d)", instanceType.String(), instanceValue.Uint())
		}
		return fmt.Sprintf("%d", instanceValue.Uint())
	case reflect.Float32, reflect.Float64:
		if math.IsNaN(instanceValue.Float()) {
			return "math.NaN()"
		}
		if reflect.TypeOf(instance) != reflect.TypeOf(float64(0.00)) && reflect.TypeOf(instance) != reflect.TypeOf(float32(0.00)) {
			return fmt.Sprintf("%s(%v)", instanceType.String(), instanceValue.Float())
		}
		return fmt.Sprintf("%v", instanceValue.Float())
	case reflect.Bool:
		if reflect.TypeOf(instance) != reflect.TypeOf(true) {
			return fmt.Sprintf("%s(%v)", instanceType.String(), instanceValue.Bool())
		}
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
