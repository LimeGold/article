package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"

	"gorm.io/gorm"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/LimeGold/article/api/v1"
	"github.com/LimeGold/article/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

// ArticleService is an article service.
type ArticleService struct {
	pb.UnimplementedArticleServer

	uc  *biz.ArticleUsecase
	log *log.Helper
}

// NewArticleService new an article service.
func NewArticleService(uc *biz.ArticleUsecase, logger log.Logger) *ArticleService {
	return &ArticleService{uc: uc, log: log.NewHelper(logger)}
}

// CreateArticle implements article.CreateArticle
func (s *ArticleService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleReply, error) {
	id, err := s.uc.Create(ctx, &biz.Article{
		Title:      req.Title,
		Content:    req.Context,
		AuthorName: req.AuthorName,
	})

	return &pb.CreateArticleReply{
		Id: id,
	}, err
}

// GetArticle implements article.GetArticle
func (s *ArticleService) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.GetArticleReply, error) {
	article, err := s.uc.Get(ctx, req.GetArticleId())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pb.ErrorArticleNotFound("article id: %d", req.GetArticleId())
		}
		return nil, err
	}

	am := &pb.ArticleMessage{
		Id:         int64(article.ID),
		Title:      article.Title,
		AuthorName: article.AuthorName,
		Content:    article.Content,
		CreatedAt:  timestamppb.New(article.CreatedAt),
		UpdatedAt:  timestamppb.New(article.UpdatedAt),
	}
	return &pb.GetArticleReply{
		Article: am,
	}, nil
}

// UpdateArticle implements article.UpdateArticle
func (s *ArticleService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleReply, error) {
	err := s.uc.Update(ctx, req.ArticleId, &biz.Article{
		Model: gorm.Model{
			ID: uint(req.ArticleId),
		},
		Title:      req.Article.Title,
		AuthorName: req.Article.AuthorName,
		Content:    req.Article.Content,
	})
	return &pb.UpdateArticleReply{}, err
}

// DeleteArticle implements article.DeleteArticle
func (s *ArticleService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleReply, error) {
	err := s.uc.Delete(ctx, req.ArticleId)
	return &pb.DeleteArticleReply{}, err
}
