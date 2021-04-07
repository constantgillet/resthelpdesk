package services

import (
	"github.com/constantgillet/resthelpdesk/config"
	"github.com/constantgillet/resthelpdesk/models"
)

func FindAllCategories() (models.Categories, error) {
	db := config.DB()
	defer db.Close()

	var categories models.Categories

	// Execute the query
	results, err := db.Query("SELECT id, name FROM categories")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {

		var category models.Category

		// for each row, scan the result into our tag composite object
		err = results.Scan(&category.Id, &category.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// push category to categories array
		categories = append(categories, category)
	}

	return categories, nil
}
