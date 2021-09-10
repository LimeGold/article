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
	data.db.AutoMigrate(&biz.Article{})
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *articleRepo) CreateArticle(ctx context.Context, g *biz.Article) (int64, error) {
	err := r.data.db.WithContext(ctx).Create(g).Error
	return int64(g.ID), err
}

func (r *articleRepo) GetArticle(ctx context.Context, id int64) (*biz.Article, error) {
	var article biz.Article
	err := r.data.db.WithContext(ctx).First(&article, "id = ?", id).Error
	return &article, err
}

func (r *articleRepo) UpdateArticle(ctx context.Context, g *biz.Article) error {
	return r.data.db.WithContext(ctx).Save(g).Error
}

func (r *articleRepo) DeleteArticle(ctx context.Context, id int64) error {
	return r.data.db.WithContext(ctx).Delete(&biz.Article{}, "id = ?", id).Error
}
