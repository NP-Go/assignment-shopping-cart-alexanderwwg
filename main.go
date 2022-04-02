/*
Features:

Main Menu:

Shopping List Application
==========================
- View Entire Shopping List
- Generate Shopping List Report
- Add Items Information
- Modify Existing Items
- Delete Item From Shopping List
- Print Current Data Fields
- Add New Category Name

*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// category list
var currentCategoryList []string

// itemInformation map
var (
	currentItemList = make(map[string](itemInformation))
)

func init() {
	var itemCategories = []string{"Household", "Food", "Drinks"}
	currentCategoryList = itemCategories
	itemList := make(map[string]itemInformation)
	itemList["Fork"] = itemInformation{0, 4, 3}
	itemList["Plates"] = itemInformation{0, 4, 3}
	itemList["Cups"] = itemInformation{0, 5, 3}
	itemList["Bread"] = itemInformation{1, 2, 2}
	itemList["Cake"] = itemInformation{1, 3, 1}
	itemList["Coke"] = itemInformation{2, 5, 2}
	itemList["Sprite"] = itemInformation{2, 5, 2}
	currentItemList = itemList
}

func main() {
	mainMenu()
}

func mainMenu() {
	fmt.Println("Shopping List Application")
	fmt.Println("=========================")
	fmt.Println("1. View entire shopping list.")
	fmt.Println("2. Generate Shopping List Report")
	fmt.Println("3. Add Items.")
	fmt.Println("4. Modify Items.")
	fmt.Println("5. Delete Item.")
	fmt.Println("6. Print Current Data.")
	fmt.Println("7. Add new Category Name")
	fmt.Println("10. Quit.")

	mainMenuOptionSelect(getIntInput())
}

func mainMenuOptionSelect(input int32) {
	switch input {
	case 1:
		viewShoppingList()
	case 2:
		generateSLReport()
	case 3:
		addItems()
	case 4:
	case 5:
	case 6:
	case 7:
	case 10:

	default:
		fmt.Println("Error, please select an option in the list.")
		mainMenu()
	}
}

// option 1
func viewShoppingList() {
	for i, v := range currentItemList {
		fmt.Print("Category: ")
		fmt.Print(currentCategoryList[v.category])
		fmt.Printf(" -  Item: %v Quantity: %v Unit Cost %v\n", i, v.quantity, v.unitCost)
	}
	fmt.Printf("Press Enter to return to the main menu.")
	fmt.Scanln()
	mainMenu()
}

// option 2
func generateSLReport() {
	fmt.Println("Generate Report")
	fmt.Println("1. Total Cost of each Category")
	fmt.Println("2. List of items by Category")
	fmt.Println("3. Main Menu")

	printCategoryInformation(getIntInput())
}

func printCategoryInformation(input int32) {
	if input == 1 || input == 2 {
		for i := range currentCategoryList {
			if input == 1 {
				var x float64
				fmt.Print(currentCategoryList[i] + ": ")
				for _, j := range currentItemList {
					if j.category == i {
						x += j.unitCost
					}
				}
				fmt.Println(x)
			}
			if input == 2 {
				for j, v := range currentItemList {
					if v.category == i {
						fmt.Print("Category: ")
						fmt.Print(currentCategoryList[v.category])
						fmt.Printf(" -  Item: %v Quantity: %v Unit Cost %v\n", j, v.quantity, v.unitCost)
					}
				}
			}
		}
		fmt.Printf("Press Enter to return to the main menu.")
		fmt.Scanln()
		mainMenu()
	} else if input == 3 {
		mainMenu()
	} else {
		fmt.Println("Error. Please try again.")
		generateSLReport()
	}
}

func addItems() {
	var name string
	var category string
	var units int
	var cost float64
AssignName:
	fmt.Println("What is the name of your item?")
	fmt.Scanln(&name)
	if name == "" {
		mainMenu()
	}
	if itemNameExists(name) {
		fmt.Println("Item already exists!")
		goto AssignName
	}
AssignCategory:
	fmt.Println("What is the category of your item?")
	fmt.Scanln(&category)
	if category == "" {
		mainMenu()
	}
	if !categoryExists(currentCategoryList, category) {
		fmt.Println("Category does not exist!\nThe categories that you have now are : ")
		printAllCategories()
		goto AssignCategory
	}
AssignUnits:
	fmt.Println("How many units?")
	fmt.Scanln(&units)
	if units == 0 {
		fmt.Println("Cannot have less than 1 unit!")
		goto AssignUnits
	}
AssignCost:
	fmt.Println("What is the cost per unit?")
	fmt.Scanln(&cost)
	if cost == 0 {
		fmt.Println("Cannot be 0!")
		goto AssignCost
	}
	newItem := itemInformation{getCategoryIndex(currentCategoryList, category), units, cost}
	currentItemList[name] = newItem
}

func printAllCategories() {
	for i := range currentCategoryList {
		fmt.Println(currentCategoryList[i])
	}
}

func itemNameExists(name string) bool {
	_, ok := currentItemList[name]
	return ok
}

func modifyQuantity() {

}

func deleteFromList() {

}

func printCurrentData() {

}

func addNewCategoryName() {

}

func getIntInput() int32 {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		fmt.Printf("Error, please use a number")
	} else {
		result := int32(input)
		return result
	}
	return 0
}
