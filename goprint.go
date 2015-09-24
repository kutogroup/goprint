package goprint

import (
	"fmt"
	"reflect"
	"strings"
)

func printStruct(t reflect.Type, v reflect.Value, space int) {
	fmt.Println("")
	for i := 0; i < t.NumField(); i++ {
		fmt.Print(strings.Repeat(" ", space), t.Field(i).Name, ":")
		value := v.Field(i)
		printValue(value, space)
		fmt.Println("")
	}
}

func printArraySlice(v reflect.Value, space int) {
	for j := 0; j < v.Len(); j++ {
		printValue(v.Index(j), space)
	}
}

func printMap(v reflect.Value, space int) {
	for _, k := range v.MapKeys() {
		printValue(k, space)
		printValue(v.MapIndex(k), space)
	}
}

func printValue(v reflect.Value, space int) {
	if !v.CanInterface() {
		fmt.Print(v)
	} else {
		printVar(v.Interface(), space)
	}
}

func printVar(i interface{}, space int) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		v = reflect.ValueOf(i).Elem()
		t = v.Type()
	}
	switch v.Kind() {
	case reflect.Array:
		printArraySlice(v, space+2)
	case reflect.Chan:
		fmt.Println("Chan")
	case reflect.Func:
		fmt.Println("Func")
	case reflect.Interface:
		fmt.Println("Interface")
	case reflect.Map:
		printMap(v, space+2)
	case reflect.Slice:
		printArraySlice(v, space+2)
	case reflect.Struct:
		printStruct(t, v, space+2)
	case reflect.UnsafePointer:
		fmt.Println("UnsafePointer")
	default:
		if v.CanInterface() {
			fmt.Print(strings.Repeat(" ", 2), v.Interface())
		} else {
			fmt.Print(strings.Repeat(" ", 2), v)
		}
	}
}

//print any variable
func P(tag string, i interface{}) {
	fmt.Println("====================", tag, "=======================")
	printVar(i, 0)
	fmt.Println("")
}

//print log by tag
func E(tag string, log string) {
	fmt.Println("====================", tag, "=======================")
	fmt.Println(log)
	fmt.Println("")
}

//print log string simply
func V(log string) {
	fmt.Println(log)
}
