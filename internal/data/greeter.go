package data

import (
	"context"

	"github.com/LimeGold/article/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

// NewArticleRepo .
func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *articleRepo) CreateArticle(ctx context.Context, g *biz.Article) error {
	return r.data.db.Create(g).Error
}

func (r *articleRepo) GetArticle(ctx context.Context, id int64) error {
	return nil
}

func (r *articleRepo) UpdateArticle(ctx context.Context, g *biz.Article) error {
	return nil
}

func (r *articleRepo) DeleteArticle(ctx context.Context, id int64) error {
	return nil
}
