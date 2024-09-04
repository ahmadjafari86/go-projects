package services

import (
	articleModels "blog/internal/modules/article/models"
	articleRepositories "blog/internal/modules/article/repositories"
	"blog/internal/modules/article/requests/articles"
	articleResponse "blog/internal/modules/article/responses"
	userResponses "blog/internal/modules/user/responses"
	"errors"
)

type ArticleService struct {
	articleRepository articleRepositories.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: articleRepositories.New(),
	}
}

func (articleService *ArticleService) GetFeaturedArticles() articleResponse.Articles {
	articles := articleService.articleRepository.List(4)
	return articleResponse.ToArticles(articles)
}

func (articleService *ArticleService) GetStoriesArticles() articleResponse.Articles {
	articles := articleService.articleRepository.List(6)
	return articleResponse.ToArticles(articles)
}

func (articleService *ArticleService) Find(id int) (articleResponse.Article, error) {
	var response articleResponse.Article
	article := articleService.articleRepository.Find(id)
	if article.ID == 0 {
		return response, errors.New("article not found")
	}
	return articleResponse.ToArticle(article), nil
}

func (articleService *ArticleService) StoreAsUser(request articles.StoreRequest, user userResponses.User) (articleResponse.Article, error) {
	var article articleModels.Article
	var response articleResponse.Article

	article.Title = request.Title
	article.Content = request.Content
	article.UserID = user.ID

	newArticle := articleService.articleRepository.Create(article)
	if newArticle.ID == 0 {
		return response, errors.New("error in creating the article")
	}
	return articleResponse.ToArticle(newArticle), nil
}
