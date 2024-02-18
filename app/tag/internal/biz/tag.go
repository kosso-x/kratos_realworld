package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// TagRepo is a Greater repo.
type TagRepo interface {
	ListTags() (tags []string, err error)
}

// TagUsecase is a Tag usecase.
type TagUsecase struct {
	repo TagRepo
	log  *log.Helper
}

// NewTagUsecase new a Tag usecase.
func NewTagUsecase(repo TagRepo, logger log.Logger) *TagUsecase {
	return &TagUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *TagUsecase) GetTags(ctx context.Context) (tags []string, err error) {
	tags, err = uc.repo.ListTags()

	if err != nil {
		return nil, err
	}

	return
}
