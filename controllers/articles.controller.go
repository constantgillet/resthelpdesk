package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/constantgillet/resthelpdesk/services"
	"github.com/gorilla/mux"
)

func GetOneArticle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	articleId, err := strconv.Atoi(vars["id"])

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app

		var resp = map[string]interface{}{
			"error": map[string]interface{}{
				"code": 422,
				"name": "Wrong article id",
			},
		}

		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Get article
	article, err := services.FindOneArticle(articleId)

	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app

		var resp = map[string]interface{}{
			"error": map[string]interface{}{
				"code": 404,
				"name": "Article not found",
			},
		}

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(resp)
		return
	}

	var resp = map[string]interface{}{
		"data": article,
	}

	json.NewEncoder(w).Encode(resp)
}

func GetAllArticles(w http.ResponseWriter, r *http.Request) {

	// Get articles
	articles, err := services.FindAllArticle()

	//If error while get articles
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
		"data": articles,
	}

	json.NewEncoder(w).Encode(resp)
}
