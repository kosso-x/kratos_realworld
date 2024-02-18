package service

import (
	"context"

	v1 "realworld/api/comment/v1"
	"realworld/app/comment/internal/biz"
)

// CommentService is a comment service.
type CommentService struct {
	v1.UnimplementedCommentServer

	uc *biz.CommentUsecase
}

// NewCommentService new a comment service.
func NewCommentService(uc *biz.CommentUsecase) *CommentService {
	return &CommentService{uc: uc}
}

func (s *CommentService) AddCommentsToAnArticle(ctx context.Context, req *v1.AddCommentsToAnArticleRequest) (res *v1.CommentReply, err error) {
	r, err := s.uc.AddCommentsToAnArticle(ctx, &biz.AddComment{
		Body: req.Comment.Body,
		Slug: req.Slug,
	})

	if err != nil {
		return nil, err
	}

	res = &v1.CommentReply{
		Comment: r,
	}

	return
}

func (s *CommentService) GetCommentsFromAnArticle(ctx context.Context, req *v1.GetCommentsRequest) (res *v1.MultipleComments, err error) {
	r, err := s.uc.GetCommentsFromAnArticle(ctx, req.Slug)
	if err != nil {
		return nil, err
	}

	res = &v1.MultipleComments{
		Comments: r,
	}

	return
}

func (s *CommentService) DeleteComment(ctx context.Context, req *v1.DeleteCommentRequest) (res *v1.DeleteCommentReply, err error) {
	if err = s.uc.DeleteComment(ctx, req.Slug, req.Id); err != nil {
		return nil, err
	}

	return
}
