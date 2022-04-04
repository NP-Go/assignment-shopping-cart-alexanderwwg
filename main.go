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
		modifyItem()
	case 5:
		deleteFromList()
	case 6:
		printCurrentData()
	case 7:
		addNewCategoryName()
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
						x += (j.unitCost * float64(j.quantity))
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
		fmt.Println("Error. No input found.")
		goto AssignName
	}
	if itemNameExists(name) {
		fmt.Println("Item already exists!")
		goto AssignName
	}
AssignCategory:
	fmt.Println("What is the category of your item?")
	fmt.Scanln(&category)
	if category == "" {
		fmt.Println("Error. No input found.")
		goto AssignCategory
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
	mainMenu()
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

// Modify Item
// First I check if I need to rename the map key-
// which will make me delete the old key and replace with the new name. But I will copy all the data required over first.
// I do not think this is the best solution at the moment. Will need to think of another way.
// Assigns all inputs to a tempItemInfo which will replace the contents of the key.
func modifyItem() {
	var name string
	var category string
	var units int
	var cost float64
	var tempItemInfo itemInformation

	fmt.Println("What item would you like to modify?")
	fmt.Scanln(&name)
	if itemNameExists(name) {
		fmt.Printf("Current item name is %v, Category is %v, Quantity is %v, Unit Cost is %v.\n", name,
			currentCategoryList[currentItemList[name].category], currentItemList[name].quantity, currentItemList[name].unitCost)
	} else {
		fmt.Println("Item does not exist.")
		modifyItem()
	}
	oldName := name
ChangeName:
	name = ""
	fmt.Println("Enter a new name. Leave empty if you do not want to change.")
	fmt.Scanln(&name)
	if name == "" {
		fmt.Println("Not renaming.")
		name = oldName
	} else if name != "" {
		if itemNameExists(name) {
			fmt.Println("Item name already exists in list.")
			goto ChangeName
		}
		currentItemList[name] = currentItemList[oldName]
		delete(currentItemList, oldName)
	}
ChangeCategory:
	fmt.Println("Enter new category. Leave empty if you do not want to change it.")
	fmt.Scanln(&category)
	if category == "" {
		fmt.Println("Not renaming.")
		tempItemInfo.category = currentItemList[name].category
	} else if !categoryExists(currentCategoryList, category) {
		fmt.Println("Category name does not exist.")
		goto ChangeCategory
	} else {
		tempItemInfo.category = getCategoryIndex(currentCategoryList, category)
	}
	fmt.Println("Enter new quantity. Leave empty if you do not wish to change it.")
	fmt.Scanln(&units)
	if units != 0 {
		tempItemInfo.quantity = units
	} else {
		tempItemInfo.quantity = currentItemList[name].quantity
	}
	fmt.Println("Enter new unit cost. Leave empty if you do not wish to change it.")
	fmt.Scanln(&cost)
	if cost != 0 {
		tempItemInfo.unitCost = cost
	} else {
		tempItemInfo.unitCost = currentItemList[name].unitCost
	}
	currentItemList[name] = tempItemInfo
	fmt.Println(currentItemList[name])
	fmt.Println(name)
	mainMenu()
}

func deleteFromList() {
	var name string
	fmt.Println("What do you want to delete?")
	fmt.Scanln(&name)
	if name == "" {
		mainMenu()
	} else {
		if itemNameExists(name) {
			delete(currentItemList, name)
			fmt.Println("Deleted " + name)
			mainMenu()
		} else {
			fmt.Println("Item not found. Nothing to delete!")
			deleteFromList()
		}
	}

}

func printCurrentData() {
	fmt.Println("Print Current Data")
	if len(currentItemList) == 0 {
		fmt.Println("No data to print!")
	} else {
		for i, v := range currentItemList {
			fmt.Printf("%v, %v \n", i, v)
		}
	}
	mainMenu()
}

func addNewCategoryName() {
	var name string
	fmt.Println("What is the new category name you would like to add?")
	fmt.Scanln(&name)
	if name == "" {
		fmt.Println("Error, no input found.")
		mainMenu()
	} else {
		if categoryExists(currentCategoryList, name) {
			fmt.Printf("Error. %v already exists.", name)
			addNewCategoryName()
		} else {
			currentCategoryList = append(currentCategoryList, name)
			fmt.Println("Added " + name + " to list of categories.")
			mainMenu()
		}
	}
}

func getIntInput() int32 {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		fmt.Println("Error, please use a number.")
	} else {
		result := int32(input)
		return result
	}
	return 0
}
