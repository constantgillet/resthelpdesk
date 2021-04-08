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
	err := db.QueryRow("SELECT id, title, content, category_id, created_at, updated_at FROM articles where id = ?", articleId).Scan(&article.Id, &article.Title, &article.Content, &article.CategoryId, &article.CreatedAt, &article.UpdatedAt)

	if err != nil {
		return article, errors.New("article not found")
	}

	return article, nil
}

type FindAllArticleOptions struct {
	Categories string
	Query      string
}

// **
// Function to get all articles
// **
func FindAllArticle(options FindAllArticleOptions) (models.Articles, error) {

	db := config.DB()
	defer db.Close()

	query := "SELECT id, title, content, category_id, created_at, updated_at FROM articles"
	var args []interface{}

	//If options is is not empty
	if (FindAllArticleOptions{}) != options {

		query = query + " WHERE"

		if options.Categories != "" {
			query = query + " category_id = ?"
			args = append(args, options.Categories)
		}

		if options.Query != "" {

			if len(args) > 0 {
				query = query + " AND"
			}

			query = query + " title LIKE ?"
			args = append(args, "%"+options.Query+"%")
		}
	}

	var articles models.Articles

	// Execute the query
	results, err := db.Query(query, args...)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {

		var article models.Article

		// for each row, scan the result into our tag composite object
		err = results.Scan(&article.Id, &article.Title, &article.Content, &article.CategoryId, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// push new article to articles array
		articles = append(articles, article)
	}

	return articles, nil
}
