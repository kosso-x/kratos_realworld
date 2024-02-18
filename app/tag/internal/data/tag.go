package data

import (
	"realworld/app/tag/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type tagRepo struct {
	data *Data
	log  *log.Helper
}

// NewTagRepo .
func NewTagRepo(data *Data, logger log.Logger) biz.TagRepo {
	return &tagRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *tagRepo) ListTags() (tags []string, err error) {
	if err = r.data.db.Table("tags").Select("name").Scan(&tags).Error; err != nil {
		return nil, err
	}

	return
}
