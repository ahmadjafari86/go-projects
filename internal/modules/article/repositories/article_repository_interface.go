package repositories

import articleModels "blog/internal/modules/article/models"

type ArticleRepositoryInterface interface {
	List(limit int) []articleModels.Article
	Find(id int) articleModels.Article
}
