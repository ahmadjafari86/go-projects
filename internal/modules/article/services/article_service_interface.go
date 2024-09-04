package services

import (
	"blog/internal/modules/article/requests/articles"
	ArticleResponse "blog/internal/modules/article/responses"
	userResponses "blog/internal/modules/user/responses"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() ArticleResponse.Articles
	GetStoriesArticles() ArticleResponse.Articles
	Find(id int) (ArticleResponse.Article, error)
	StoreAsUser(request articles.StoreRequest, user userResponses.User) (ArticleResponse.Article, error)
}
