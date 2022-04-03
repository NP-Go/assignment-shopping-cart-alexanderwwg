package main

import (
	"reflect"
)

func categoryExists(slice []string, index string) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("Invalid Data Type!")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == index {
			return true
		}
	}
	return false
}

func getCategoryIndex(slice []string, index string) int {
	s := reflect.ValueOf(slice)

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == index {
			return i
		}
	}
	return 0
}
