package repositories

import (
	ArticleModels "blog/internal/modules/article/models"
	"blog/pkg/database"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func New() *ArticleRepository {
	return &ArticleRepository{
		DB: database.Connection(),
	}
}

func (articleRepository *ArticleRepository) List(limit int) []ArticleModels.Article {
	var articles []ArticleModels.Article
	articleRepository.DB.Limit(limit).Joins("User").Order("rand()").Find(&articles)
	return articles
}

func (articleRepository *ArticleRepository) Find(id int) ArticleModels.Article {
	var article ArticleModels.Article
	articleRepository.DB.Joins("User").First(&article, id)
	return article
}
