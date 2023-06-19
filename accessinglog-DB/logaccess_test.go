package main

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestGetTopNMenuItems(t *testing.T) {
	menuItemCounts := map[string]int{
		"Item1": 10,
		"Item2": 5,
		"Item3": 8,
		"Item4": 3,
		"Item5": 12,
	}

	expected := []string{"Item5", "Item1", "Item3"}

	topNMenuItems := getTopNMenuItems(menuItemCounts, 3)
	fmt.Println(topNMenuItems)

	assert.Equal(t, expected, topNMenuItems, "The top N menu items are not as expected")
}

func TestIntegration(t *testing.T) {
	// Connect to the MySQL database
	db, err := sql.Open("mysql", "sammy:Password@1@tcp(localhost:3306)/mydb")
	if err != nil {
		t.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT eater_id, foodmenu_id FROM log")
	if err != nil {
		t.Fatal("Failed to retrieve log from the database:", err)
	}
	defer rows.Close()

	menuItemCounts := make(map[string]int)

	for rows.Next() {
		var eaterID, foodmenuID string
		err := rows.Scan(&eaterID, &foodmenuID)
		if err != nil {
			t.Fatal("Failed to scan log row:", err)
		}

		menuItemCounts[foodmenuID]++
	}

	top3MenuItems := getTopNMenuItems(menuItemCounts, 3)

	// Ensure the top 3 menu items are not empty
	assert.NotEmpty(t, top3MenuItems, "Top 3 menu items should not be empty")
}
