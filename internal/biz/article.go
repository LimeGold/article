package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title      string
	AuthorName string
	Content    string
}

type ArticleRepo interface {
	CreateArticle(context.Context, *Article) (int64, error)
	GetArticle(context.Context, int64) (*Article, error)
	UpdateArticle(context.Context, *Article) error
	DeleteArticle(context.Context, int64) error
}

type ArticleUsecase struct {
	repo ArticleRepo
	log  *log.Helper
}

func NewArticleUsecase(repo ArticleRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ArticleUsecase) Create(ctx context.Context, g *Article) (int64, error) {
	return uc.repo.CreateArticle(ctx, g)
}

func (uc *ArticleUsecase) Update(ctx context.Context, g *Article) error {
	return uc.repo.UpdateArticle(ctx, g)
}

func (uc *ArticleUsecase) Get(ctx context.Context, id int64) (*Article, error) {
	return uc.repo.GetArticle(ctx, id)
}

func (uc *ArticleUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteArticle(ctx, id)
}
