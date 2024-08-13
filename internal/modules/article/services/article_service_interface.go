package services

import ArticleModels "blog/internal/modules/article/models"

type ArticleServiceInterface interface {
	GetFeaturedArticles() []ArticleModels.Article
	GetStoriesArticles() []ArticleModels.Article
}
