package biz

import (
	"context"
	v1 "realworld/api/article/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gogf/gf/v2/util/gconv"
)

// ArticleRepo is a Greater repo.
type ArticleRepo interface {
	SearchArticles(search *ListReq) (res []*ArticleReply, err error)
	WithTagsFavoritesCount(atcs []*ArticleReply, author *User) (err error)
	FeedArticles(search *FeedReq) (res []*ArticleReply, err error)
	SlugArticle(slug string) (res []*ArticleReply, err error)
	CreateArticle(article *CreateReq) (res *ArticleReply, err error)
	UpdateArticle(article *CreateReq, slug string) (res *ArticleReply, err error)
	DeleteArticle(slug string) (err error)
	FavoriteArticle(slug string) (res *ArticleReply, err error)
	UnfavoriteArticle(slug string) (res *ArticleReply, err error)
}

// ArticleUsecase is a Article usecase.
type ArticleUsecase struct {
	repo ArticleRepo
	log  *log.Helper
}

// NewArticleUsecase new a Article usecase.
func NewArticleUsecase(repo ArticleRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ArticleUsecase) ListArticles(ctx context.Context, req *ListReq) (res []*v1.ArticlesReply_Articles, err error) {
	r, err := uc.repo.SearchArticles(req)
	if err != nil {
		return nil, err
	}

	if err = gconv.Scan(r, &res); err != nil {
		return nil, err
	}

	return
}

func (uc *ArticleUsecase) FeedArticles(ctx context.Context, req *FeedReq) (res []*v1.ArticlesReply_Articles, err error) {
	r, err := uc.repo.FeedArticles(req)
	if err != nil {
		return nil, err
	}

	if err = gconv.Scan(r, &res); err != nil {
		return nil, err
	}

	return
}

func (uc *ArticleUsecase) GetArticle(ctx context.Context, slug string) (res []*v1.ArticlesReply_Articles, err error) {
	r, err := uc.repo.SlugArticle(slug)

	if err != nil {
		return nil, err
	}

	if err = gconv.Scan(r, &res); err != nil {
		return nil, err
	}

	return
}

func (uc *ArticleUsecase) CreateArticle(ctx context.Context, req *v1.CreateArticleRequest_CreateArticle) (res *v1.ArticlesReply_Articles, err error) {
	r, err := uc.repo.CreateArticle(&CreateReq{
		Title:       req.Title,
		Description: req.Description,
		Body:        req.Body,
		TagList:     req.TagList,
	})
	if err != nil {
		return nil, err
	}

	if err = gconv.Scan(r, &res); err != nil {
		return nil, err
	}

	return
}

func (uc *ArticleUsecase) UpdateArticle(ctx context.Context, req *v1.UpdateArticleRequest_UpdateArticle, slug string) (res *v1.ArticlesReply_Articles, err error) {
	r, err := uc.repo.UpdateArticle(&CreateReq{
		Title:       req.Title,
		Description: req.Description,
		Body:        req.Body,
		TagList:     req.TagList,
	}, slug)

	if err != nil {
		return nil, err
	}

	if err = gconv.Scan(r, &res); err != nil {
		return nil, err
	}

	return
}

func (uc *ArticleUsecase) DeleteArticle(ctx context.Context, slug string) (err error) {
	if err = uc.repo.DeleteArticle(slug); err != nil {
		return err
	}

	return
}

func (uc *ArticleUsecase) FavoriteArticle(ctx context.Context, slug string) (res *v1.ArticlesReply_Articles, err error) {
	r, err := uc.repo.FavoriteArticle(slug)
	if err != nil {
		return nil, err
	}

	if err = gconv.Scan(r, &res); err != nil {
		return nil, err
	}

	return
}

func (uc *ArticleUsecase) UnfavoriteArticle(ctx context.Context, slug string) (res *v1.ArticlesReply_Articles, err error) {
	r, err := uc.repo.UnfavoriteArticle(slug)
	if err != nil {
		return nil, err
	}

	if err = gconv.Scan(r, &res); err != nil {
		return nil, err
	}

	return
}
