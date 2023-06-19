package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type MenuItemCount struct {
	MenuItem string
	Count    int
}

type menuItem struct {
	EaterID string
	FoodID  string
}

// ItemsK represents a key-value pair
type ItemsK struct {
	Key   string
	Count int
}

func main() {
	// Read the log file
	filePath := "log.txt"
	logFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer logFile.Close()

	// Map to store the count of each menu item
	menuItemCounts := make(map[string]int)
	menuItems := []menuItem{}

	// Scan the log file line by line
	scanner := bufio.NewScanner(logFile)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) != 2 {
			log.Println("Invalid log entry:", line)
			continue
		}

		eaterID := fields[0]
		foodmenuID := fields[1]

		menuItemCounts[foodmenuID]++

		menuItems = append(menuItems, menuItem{eaterID, foodmenuID})

	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error occurred while scanning log file:", err)
	}

	// fmt.Println(menuItems)

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
	top3MenuItems := findTopNMenuItems(menuItemCounts, 3)

	// Print the top 3 menu items
	fmt.Println("Top 3 Menu Items Consumed:")
	for _, menuItem := range top3MenuItems {
		fmt.Println(menuItem)
	}
}

// Function to get the top N menu items based on their counts
func findTopNMenuItems(menuItemCounts map[string]int, N int) []string {
	t := 0
	itemTemp := []ItemsK{}
	for key, count := range menuItemCounts {
		itemTemp = append(itemTemp, ItemsK{key, count})
	}

	for i := 0; i < len(itemTemp)-1; i++ {
		t = i
		for j := i + 1; j < len(itemTemp); j++ {
			if itemTemp[t].Count < itemTemp[j].Count {
				t = j
			}
		}
		// fmt.Println(i, t)
		itemTemp[i], itemTemp[t] = itemTemp[t], itemTemp[i]
	}
	// Get the top N menu items
	topNMenuItems := make([]string, N)

	for i := 0; i < N; i++ {
		topNMenuItems[i] = itemTemp[i].Key
	}

	// fmt.Println(topNMenuItems)
	return topNMenuItems
}
