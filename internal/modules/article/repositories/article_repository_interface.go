package repositories

import ArticleModels "blog/internal/modules/article/models"

type ArticleRepositoryInterface interface {
	List(limit int) []ArticleModels.Article
}
