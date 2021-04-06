package services

import (
	"errors"

	"github.com/constantgillet/resthelpdesk/config"
	"github.com/constantgillet/resthelpdesk/models"
)

// **
// Function to get one article
// **
func FindOneArticle(articleId int) (models.Article, error) {

	db := config.DB()
	defer db.Close()

	var article models.Article

	// Execute the query
	err := db.QueryRow("SELECT id, title, content, created_at, updated_at FROM articles where id = ?", articleId).Scan(&article.Id, &article.Title, &article.Content, &article.CreatedAt, &article.UpdatedAt)

	if err != nil {
		return article, errors.New("article not found")
	}

	return article, nil
}

// **
// Function to get all articles
// **
func FindAllArticle() (models.Articles, error) {

	db := config.DB()
	defer db.Close()

	var articles models.Articles

	// Execute the query
	results, err := db.Query("SELECT id, title, content, created_at, updated_at FROM articles")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {

		var article models.Article

		// for each row, scan the result into our tag composite object
		err = results.Scan(&article.Id, &article.Title, &article.Content, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// push new article to articles array
		articles = append(articles, article)
	}

	return articles, nil
}
