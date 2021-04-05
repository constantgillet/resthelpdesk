package services

import (
	"errors"

	"github.com/constantgillet/resthelpdesk/config"
	"github.com/constantgillet/resthelpdesk/models"
)

func FindOneArticle(articleId int) (models.Article, error) {

	db := config.DB()

	var article models.Article

	// Execute the query
	err := db.QueryRow("SELECT id, title, content, created_at FROM articles where id = ?", articleId).Scan(&article.Id, &article.Title, &article.Content, &article.CreatedAt)

	defer db.Close()

	if err != nil {
		return article, errors.New("article not found")
	}

	return article, nil

}
