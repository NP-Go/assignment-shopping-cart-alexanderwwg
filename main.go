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
	"fmt"
)

type itemInformation struct {
	category int
	quantity int
	unitCost float64
}

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
	fmt.Println("3.Add Items.")
	fmt.Println("4.Modify Items.")
	fmt.Println("5.Delete Item.")
	fmt.Println("6.Print Current Data.")
	fmt.Println("7.Add new Category Name")
	fmt.Println("Input 0 to quit.")

	optionSelect(getIntInput())
}

func optionSelect(input int) {
	switch input {
	// end program
	case 0:
	// if not do everything else
	case 1:
		viewShoppingList()
	case 2:
		generateSLReport()
	case 3:
	case 4:
	case 5:
	case 6:
	case 7:
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

func printCategoryInformation(input int) {
	if input == 1 || input == 2 {
		for i := range currentCategoryList {
			if input == 1 {
				var x float64
				fmt.Print(currentCategoryList[i] + ": ")
				for _, v := range currentItemList {
					if v.category == i {
						x += v.unitCost
					}
				}
				fmt.Println(x)
			}
			if input == 2 {
				for k, v := range currentItemList {
					if v.category == i {
						fmt.Print("Category: ")
						fmt.Print(currentCategoryList[v.category])
						fmt.Printf(" -  Item: %v Quantity: %v Unit Cost %v\n", k, v.quantity, v.unitCost)
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

func addItemInformation() {

}

func modifyQuantity() {

}

func deleteFromList() {

}

func printCurrentData() {

}

func addNewCategoryName() {

}

func getIntInput() int {
	var input int
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error: Please use numbers only.")
		getIntInput()
	}
	return input
}
