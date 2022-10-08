package main

import "errors"

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `jsonn:"content"`
}

var articleList = []Article{
	{ID: 1, Title: "Article 1", Content: "Article 1 Body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 Body"},
}

func getAllArticles() []Article {
	return articleList
}

func getArticleByID(articleId int) (*Article, error) {
	for _, article := range articleList {
		if article.ID == articleId {
			return &article, nil
		}
	}
	return nil, errors.New("Article not found")
}
