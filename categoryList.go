package main

import (
	"fmt"
	"reflect"
)

func addCategoryToSlice(categoryName string) {

	// check if name exists in slice
	if categoryExists(currentCategoryList, categoryName) {
		fmt.Println("True")
	} else {
		fmt.Println("Adding " + categoryName + " to list of categories.")
		currentCategoryList = append(currentCategoryList, categoryName)
	}

}

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
