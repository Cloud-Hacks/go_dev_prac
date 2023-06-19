package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MenuItemCount struct {
	MenuItem string
	Count    int
}

// ItemsK represents a key-value pair
type ItemsK struct {
	Key   string
	Count int
}

type menuItem struct {
	EaterID string
	FoodID  string
}

// var menuItemsID = make(map[string]string, 10)

func main() {
	// Connect to the MySQL database
	db, err := sql.Open("mysql", "sammy:Password@1@tcp(localhost:3306)/mydb")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	/* 	menuItemsID := map[string]string{
	   		"45ER": "45",
	   		"76TR": "26",
	   		"85FG": "84",
	   	}

	   	for _, j := range menuItemsID {
	   		j, _ := strconv.Atoi(j)
	   		fmt.Println(j)
	   	} */

	// Read the log from the database
	rows, err := db.Query("SELECT eater_id, foodmenu_id FROM log")
	if err != nil {
		log.Fatal("Failed to retrieve log from the database:", err)
	}
	defer rows.Close()

	// Map to store the count of each menu item
	menuItemCounts := make(map[string]int)
	menuItems := []menuItem{}

	// Iterate over each row in the log
	for rows.Next() {
		var eaterID, foodmenuID string
		err := rows.Scan(&eaterID, &foodmenuID)
		if err != nil {
			log.Println("Failed to scan log row:", err)
			continue
		}

		// menuItems[eaterID] = foodmenuID
		// fmt.Println(eaterID, foodmenuID)

		// Increment the count for the menu item
		menuItemCounts[foodmenuID]++

		menuItems = append(menuItems, menuItem{eaterID, foodmenuID})

	}
	if err := rows.Err(); err != nil {
		log.Fatal("Error occurred while iterating over log rows:", err)
	}

	// Check for duplicate eater_id-foodmenu_id combination
	for i := 0; i < len(menuItems); i++ {
		for j := i + 1; j < len(menuItems); j++ {
			if menuItems[i].EaterID == menuItems[j].EaterID {
				if menuItems[i].FoodID == menuItems[j].FoodID {
					log.Fatalf("Error: Duplicate entry for eater_id %s and foodmenu_id %s\n", menuItems[i].EaterID, menuItems[i].FoodID)
				}
			}
		}
	}

	// Get the top 3 menu items consumed
	top3MenuItems := getTopNMenuItems(menuItemCounts, 3)

	// topMenu := make(map[string]int)

	// itemK := make([]int, 3)

	// k := 0
	// t := 0

	// var str string

	// fmt.Println(menuItems)

	/* 	for i, n := range menuItems {
		k = 1
		for m, j := range menuItems {
			if menuItems[i] == j {
				if k > 1 {
					fmt.Println(menuItems[i], j, i, m, n)
					if i == m {
						fmt.Println("Duplicate items", i)
						break
					}
				}
				k++
			}
		}
	} */

	/* 	for _, j := range menuItems {
		for i, _ := range menuItems {
			if j == menuItems[i] {
				topMenu[j]++
			}
		}
	} */

	/* 	menuId := make([]int, 0, len(topMenu))

	   	topMenu["34"] = 8
	   	for i, j := range topMenu {
	   		if j > 1 {
	   			topMenu[i] = j / 2
	   		}
	   		menuId = append(menuId, topMenu[i])
	   		// menuId[t] = topMenu[i]
	   		t++
	   	}
	   	fmt.Println(menuId) */
	// fmt.Println(topMenu)

	// sorting based on no. of counts(j)
	/* 	for _, j := range topMenu {
		// tmp := m
		for i, _ := range topMenu {
			// str = i
			if j < topMenu[i] {
				// if i == tmp {

				// }
				j = topMenu[i]
				// fmt.Println(i, j)
			}
		}
		// topMenu[str] = j
		if k < 3 {
			itemK[k] = j
			k++
		} else {
			break
		}
	} */

	// sort.Sort(sort.Reverse(sort.IntSlice(menuId)))
	// fmt.Println(menuId)

	/* 	for k := 0; k < 3; k++ {
		for i, _ := range topMenu {
			if menuId[k] == topMenu[i] {
				// fmt.Println(i, k)
				fmt.Println(i)
				break
			}
		}
	} */

	// fmt.Println(topMenu, menuItemCounts)

	// Print the top 3 menu items
	fmt.Println("Top 3 Menu Items Consumed:")
	for _, menuItem := range top3MenuItems {
		fmt.Println(menuItem)
	}
}

// Function to get the top N menu items based on their counts
func getTopNMenuItems(menuItemCounts map[string]int, N int) []string {
	/* 	type MenuItemCount struct {
	   		MenuItem string
	   		Count    int
	   	}

	   	// Create a slice of MenuItemCount structs
	   	// menuItems := make([]MenuItemCount, len(menuItemCounts))
	   	menuItems := []MenuItemCount{}

	   	// Populate the slice with menu items and their counts
	   	for menuItem, count := range menuItemCounts {
	   		menuItems = append(menuItems, MenuItemCount{menuItem, count})
	   	}

	   	// Sort the menu items based on their counts (descending order)
	   	sort.Slice(menuItems, func(i, j int) bool {
	   		return menuItems[i].Count > menuItems[j].Count
	   	}) */

	t := 0
	itemTemp := []ItemsK{}
	for key, count := range menuItemCounts {
		itemTemp = append(itemTemp, ItemsK{key, count})
	}

	for i := 0; i < len(itemTemp)-1; i++ {
		t = i
		for j := i + 1; j < len(itemTemp); j++ {
			if itemTemp[t].Count < itemTemp[j].Count {
				// j = i
				t = j
				// i = j
			}
		}
		// fmt.Println(i, t)
		itemTemp[i], itemTemp[t] = itemTemp[t], itemTemp[i]
	}

	// fmt.Println(itemTemp)

	// Get the top N menu items
	topNMenuItems := make([]string, N)

	for i := 0; i < N; i++ {
		topNMenuItems[i] = itemTemp[i].Key
	}

	/* 	for i := 0; i < N && i < len(menuItems); i++ {
		topNMenuItems = append(topNMenuItems, menuItems[i].MenuItem)
	} */

	// fmt.Println(topNMenuItems)
	return topNMenuItems
}
