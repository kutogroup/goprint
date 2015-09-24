package goprint

import (
	"fmt"
	"reflect"
	"strings"
)

func printStruct(t reflect.Type, v reflect.Value) {
	fmt.Println("")
	for i := 0; i < t.NumField(); i++ {
		fmt.Print(strings.Repeat(" ", 0), t.Field(i).Name, ":")
		value := v.Field(i)
		if !value.CanInterface() {
			fmt.Println(" is nil")
		} else {
			printVar(value.Interface())
			fmt.Println("")
		}
	}
}

func printArraySlice(v reflect.Value) {
	for j := 0; j < v.Len(); j++ {
		printVar(v.Index(j).Interface())
	}
}

func printMap(v reflect.Value) {
	for _, k := range v.MapKeys() {
		printVar(k.Interface())
		printVar(v.MapIndex(k).Interface())
	}
}

func printVar(i interface{}) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		v = reflect.ValueOf(i).Elem()
		t = v.Type()
	}
	switch v.Kind() {
	case reflect.Array:
		printArraySlice(v)
	case reflect.Chan:
		fmt.Println("Chan")
	case reflect.Func:
		fmt.Println("Func")
	case reflect.Interface:
		fmt.Println("Interface")
	case reflect.Map:
		printMap(v)
	case reflect.Slice:
		printArraySlice(v)
	case reflect.Struct:
		printStruct(t, v)
	case reflect.UnsafePointer:
		fmt.Println("UnsafePointer")
	default:
		fmt.Print(strings.Repeat(" ", 2), v.Interface())
	}
}

//输出任意变量的值
func P(i interface{}) {
	fmt.Println("====================================================")
	printVar(i)
}

//输出带标签的字符串值
func E(tag string, log string) {
	fmt.Println("====================", tag, "=======================")
	fmt.Println(log)
	fmt.Println("")
}

//输出原始字符串值
func V(log string) {
	fmt.Println(log)
}
