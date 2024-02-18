package biz

import (
	"context"
	v1 "realworld/api/comment/v1"

	"github.com/go-kratos/kratos/v2/log"
)

// CommentRepo is a Greater repo.
type CommentRepo interface {
	GetArticle(slug string) (article *Article, err error)
	CreateComment(comment *Comment) (res *v1.CommentReply_Comment, err error)
	ListComments(article_id int64) (res []*v1.MultipleComments_Comments, err error)
	DeleteComment(slug string, comment_id int64) (err error)
}

// CommentUsecase is a Comment usecase.
type CommentUsecase struct {
	repo CommentRepo
	log  *log.Helper
}

// NewCommentUsecase new a Comment usecase.
func NewCommentUsecase(repo CommentRepo, logger log.Logger) *CommentUsecase {
	return &CommentUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CommentUsecase) AddCommentsToAnArticle(ctx context.Context, add *AddComment) (res *v1.CommentReply_Comment, err error) {
	article, err := uc.repo.GetArticle(add.Slug)
	if err != nil {
		return nil, err
	}

	res, err = uc.repo.CreateComment(&Comment{
		Body:      add.Body,
		ArticleId: int64(article.ID),
		UserId:    2, // rwmiddleware.UserClaims.Id,
	})

	return
}

func (uc *CommentUsecase) GetCommentsFromAnArticle(ctx context.Context, slug string) (res []*v1.MultipleComments_Comments, err error) {
	article, err := uc.repo.GetArticle(slug)
	if err != nil {
		return nil, err
	}

	res, err = uc.repo.ListComments(int64(article.ID))
	if err != nil {
		return nil, err
	}

	return
}

func (uc *CommentUsecase) DeleteComment(ctx context.Context, slug string, comment_id int64) (err error) {
	if err := uc.repo.DeleteComment(slug, comment_id); err != nil {
		return err
	}

	return
}
