package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/constantgillet/resthelpdesk/services"
)

func GetAllCategories(w http.ResponseWriter, r *http.Request) {

	// Get categories
	categories, err := services.FindAllCategories()

	//If error while get categories
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app

		var resp = map[string]interface{}{
			"error": map[string]interface{}{
				"code": 500,
				"name": "Internal server error",
			},
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resp)
		return
	}

	var resp = map[string]interface{}{
		"data": categories,
	}

	json.NewEncoder(w).Encode(resp)
}
