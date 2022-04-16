package repository

import "blog/model"

func ArticleList() ([]*model.Article, error) {
	query := `SELECT * FROM  article`

	var article []*model.Article
	if err := db.Select(&article, query); err !=nil {
		return nil, err
	}

	return article, nil
}