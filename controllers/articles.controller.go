package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/constantgillet/resthelpdesk/config"
	"github.com/constantgillet/resthelpdesk/services"
	"github.com/gorilla/mux"
)

func GetAllArticles() {

	// db := config.DB()

	// // Execute the query
	// results, err := db.Query("SELECT id, title, content, created_at FROM articles")
	// if err != nil {
	//     panic(err.Error()) // proper error handling instead of panic in your app
	// }

	// var articles models.Articles

	// for results.Next() {
	//     var article models.Article
	//     // for each row, scan the result into our tag composite object
	//     err = results.Scan(&article.Id, &article.Title, &article.Content, &article.CreatedAt)
	//     if err != nil {
	//         panic(err.Error()) // proper error handling instead of panic in your app
	//     }

	// }

	// defer db.Close()

}

func GetOneArticle(w http.ResponseWriter, r *http.Request) {

	db := config.DB()

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

	// Execute the query
	article, err := services.FindOneArticle(articleId)

	defer db.Close()

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
