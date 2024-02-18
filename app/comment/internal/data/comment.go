package data

import (
	v1 "realworld/api/comment/v1"
	"realworld/app/comment/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type commentRepo struct {
	data *Data
	log  *log.Helper
}

// NewCommentRepo .
func NewCommentRepo(data *Data, logger log.Logger) biz.CommentRepo {
	return &commentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *commentRepo) GetArticle(slug string) (article *biz.Article, err error) {
	if err = r.data.db.Table("articles").Where("slug = ?", slug).Find(&article).Error; err != nil {
		return nil, err
	}

	return
}

func (r *commentRepo) CreateComment(comment *biz.Comment) (res *v1.CommentReply_Comment, err error) {
	var c *biz.Comment
	if err = r.data.db.Create(comment).Scan(&c).Error; err != nil {
		return nil, err
	}

	res = &v1.CommentReply_Comment{
		Id:        int64(c.ID),
		CreatedAt: timestamppb.New(c.CreatedAt),
		UpdatedAt: timestamppb.New(c.UpdatedAt),
		Body:      c.Body,
	}

	if err = r.data.db.Table("users").Where("id = ?", comment.UserId).Find(&res.Author).Error; err != nil {
		return nil, err
	}

	return
}

func (r *commentRepo) ListComments(article_id int64) (res []*v1.MultipleComments_Comments, err error) {
	var (
		cs      []*biz.Comment
		authors = make(map[int64]*v1.MultipleComments_Author)
	)

	if err = r.data.db.Table("comments").Where("article_id = ?", article_id).Find(&cs).Error; err != nil {
		return nil, err
	}

	for _, c := range cs {
		author := authors[c.UserId]
		if author == nil {
			if err = r.data.db.Table("users").Where("id = ?", c.UserId).Find(&author).Error; err != nil {
				return nil, err
			}
			authors[c.UserId] = author
			author = authors[c.UserId]
		}

		res = append(res, &v1.MultipleComments_Comments{
			Id:        uint32(c.ID),
			CreatedAt: timestamppb.New(c.CreatedAt),
			UpdatedAt: timestamppb.New(c.UpdatedAt),
			Body:      c.Body,
			Author:    author,
		})
	}

	return
}

func (r *commentRepo) DeleteComment(slug string, comment_id int64) (err error) {
	if err = r.data.db.Table("comments").Where("id = ?", comment_id).Delete(&biz.Comment{}).Error; err != nil {
		return err
	}

	return
}
