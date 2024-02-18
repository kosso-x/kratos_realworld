package service

import (
	"context"
	v1 "realworld/api/article/v1"
	"realworld/app/article/internal/biz"

	"github.com/gogf/gf/v2/util/gconv"
)

// ArticleService is a article service.
type ArticleService struct {
	v1.UnimplementedArticleServer

	uc *biz.ArticleUsecase
}

// NewArticleService new a article service.
func NewArticleService(uc *biz.ArticleUsecase) *ArticleService {
	return &ArticleService{uc: uc}
}

func (s *ArticleService) ListArticles(ctx context.Context, req *v1.ListArticlesRequest) (res *v1.ArticlesReply, err error) {
	r, err := s.uc.ListArticles(ctx, &biz.ListReq{
		Tag:       req.Tag,
		Author:    req.Author,
		Favorited: req.Favorited,
		Limit:     req.Limit,
		Offset:    req.Offset,
	})

	if err != nil {
		return nil, err
	}

	res = &v1.ArticlesReply{
		Articles:      r,
		ArticlesCount: uint32(len(r)),
	}

	return
}

func (s *ArticleService) FeedArticles(ctx context.Context, req *v1.FeedArticlesRequest) (res *v1.ArticlesReply, err error) {
	r, err := s.uc.FeedArticles(ctx, &biz.FeedReq{
		Limit:  req.Limit,
		Offset: req.Offset,
	})

	if err != nil {
		return nil, err
	}

	res = &v1.ArticlesReply{
		Articles:      r,
		ArticlesCount: uint32(len(r)),
	}
	return
}

func (s *ArticleService) GetArticle(ctx context.Context, req *v1.GetArticleRequest) (res *v1.ArticleReply, err error) {
	r, err := s.uc.GetArticle(ctx, req.Slug)

	if err != nil {
		return nil, err
	}

	if err = gconv.Struct(r[0], &res); err != nil {
		return nil, err
	}

	return
}

func (s *ArticleService) CreateArticle(ctx context.Context, req *v1.CreateArticleRequest) (res *v1.ArticleReply, err error) {
	r, err := s.uc.CreateArticle(ctx, req.Article)

	if err != nil {
		return nil, err
	}

	if err = gconv.Struct(r, &res); err != nil {
		return nil, err
	}

	return
}

func (s *ArticleService) UpdateArticle(ctx context.Context, req *v1.UpdateArticleRequest) (res *v1.ArticleReply, err error) {
	r, err := s.uc.UpdateArticle(ctx, req.Article, req.Slug)
	if err != nil {
		return nil, err
	}

	if err = gconv.Struct(r, &res); err != nil {
		return nil, err
	}

	return
}

func (s *ArticleService) DeleteArticle(ctx context.Context, req *v1.DeleteArticleRequest) (res *v1.DeleteArticleReply, err error) {
	if err = s.uc.DeleteArticle(ctx, req.Slug); err != nil {
		return nil, err
	}

	return
}

func (s *ArticleService) FavoriteArticle(ctx context.Context, req *v1.FavoriteArticleRequest) (res *v1.ArticleReply, err error) {
	r, err := s.uc.FavoriteArticle(ctx, req.Slug)

	if err != nil {
		return nil, err
	}

	if err = gconv.Struct(r, &res); err != nil {
		return nil, err
	}

	return
}

func (s *ArticleService) UnfavoriteArticle(ctx context.Context, req *v1.UnfavoriteArticleRequest) (res *v1.ArticleReply, err error) {
	r, err := s.uc.UnfavoriteArticle(ctx, req.Slug)

	if err != nil {
		return nil, err
	}

	if err = gconv.Struct(r, &res); err != nil {
		return nil, err
	}

	return
}
