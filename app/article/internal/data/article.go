package data

import (
	"fmt"
	"realworld/app/article/internal/biz"
	cryptomd5 "realworld/package/cryptoMd5"
	rwmiddleware "realworld/package/rwMiddleware"

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

func (r *articleRepo) SearchArticles(search *biz.ListReq) (res []*biz.ArticleReply, err error) {
	/*
		1. 通过 tag                    找 tag_id,  在 article_tags      找 article_ids
		2. 通过 favorited              找 user_id, 在 article_favorites 找 article_ids
		3. 通过 author 和 article_ids, 找 user_id, 在 articles          找 articles
	*/

	var (
		tag_ids       []int64
		t_article_ids []int64
		favorite_ids  []int64
		f_article_ids []int64
		author        *biz.User
		article_ids   []int64
	)

	if err = r.data.db.Table("tags").Select("id").Where(&biz.Tag{
		Name: search.Tag,
	}).Find(&tag_ids).Error; err != nil {
		return nil, err
	}

	if err = r.data.db.Table("article_tags").Select("article_id").
		Where("tag_id in (?)", tag_ids).Find(&t_article_ids).Error; err != nil {
		return nil, err
	}

	if err = r.data.db.Table("users").Select("id").Where(&biz.User{
		UserName: search.Favorited,
	}).Find(&favorite_ids).Error; err != nil {
		return nil, err
	}

	if err = r.data.db.Table("article_favorites").Select("article_id").
		Where("user_id in (?)", favorite_ids).Find(&f_article_ids).Error; err != nil {
		return nil, err
	}

	if err = r.data.db.Table("users").Where(&biz.User{
		UserName: search.Author,
	}).Find(&author).Error; err != nil {
		return nil, err
	}

	article_ids = append(t_article_ids, f_article_ids...)

	if article_ids != nil {
		if err = r.data.db.Table("articles").Limit(int(search.Limit)).Offset(int(search.Offset)).Where("id in (?)", article_ids).Where(&biz.Article{
			UserId: author.Id,
		}).Find(&res).Error; err != nil {
			return nil, err
		}
	} else {
		if err = r.data.db.Table("articles").Limit(int(search.Limit)).Offset(int(search.Offset)).Where(&biz.Article{
			UserId: author.Id,
		}).Find(&res).Error; err != nil {
			return nil, err
		}
	}

	if err = r.WithTagsFavoritesCount(res, author); err != nil {
		return nil, err
	}

	return
}

func (r *articleRepo) WithTagsFavoritesCount(atcs []*biz.ArticleReply, author *biz.User) (err error) {
	var user_ids []int64
	for _, atc := range atcs {
		rows, _ := r.data.db.Raw("select t.name from tags t left join article_tags ats on t.id = ats.tag_id where ats.article_id = ?", atc.ID).Rows()
		if rows == nil {
			continue
		}
		defer rows.Close()
		for rows.Next() {
			var text string
			rows.Scan(&text)
			atc.TagList = append(atc.TagList, text)
		}

		if err = r.data.db.Table("article_favorites").Select("user_id").Where(&biz.ArticleFavorite{
			ArticleId: int64(atc.ID),
		}).Find(&user_ids).Error; err != nil {
			return err
		}

		atc.FavoritesCount = len(user_ids)
		_, atc.Favorited = r.FindItem(user_ids, rwmiddleware.UserClaims.Id)
		// _, atc.Favorited = r.FindItem(user_ids, 2)
		atc.Author = author
	}

	return
}

func (r *articleRepo) FeedArticles(search *biz.FeedReq) (res []*biz.ArticleReply, err error) {
	var (
		follow_ids []int64
		author     *biz.User
	)

	if err = r.data.db.Table("follows").Select("follow_id").Where("user_id = ?", 2).Find(&follow_ids).Error; err != nil {
		return nil, err
	}

	if err = r.data.db.Table("articles").Limit(int(search.Limit)).Offset(int(search.Offset)).Where("user_id in (?)", follow_ids).Find(&res).Error; err != nil {
		return nil, err
	}

	if err = r.data.db.Table("users").Where("id = ?", 2).Find(&author).Error; err != nil {
		return nil, err
	}

	if err = r.WithTagsFavoritesCount(res, author); err != nil {
		return nil, err
	}

	return
}

func (r *articleRepo) SlugArticle(slug string) (res []*biz.ArticleReply, err error) {
	var author *biz.User

	if err = r.data.db.Table("articles").Where("slug = ?", slug).Find(&res).Error; err != nil {
		return nil, err
	}

	if err = r.data.db.Table("users").Where("id = ?", 2).Find(&author).Error; err != nil {
		return nil, err
	}

	if err = r.WithTagsFavoritesCount(res, author); err != nil {
		return nil, err
	}

	return
}

func (r *articleRepo) CreateArticle(article *biz.CreateReq) (res *biz.ArticleReply, err error) {
	var author *biz.User
	if err = r.data.db.Create(&biz.Article{
		Title:       article.Title,
		Description: article.Description,
		Body:        article.Body,
		UserId:      rwmiddleware.UserClaims.Id,
		Slug:        cryptomd5.Md5To16(fmt.Sprintf("%v", article)),
	}).Scan(&res).Error; err != nil {
		return nil, err
	}

	if err = r.data.db.Table("users").Where("id = ?", 2).Find(&author).Error; err != nil {
		return nil, err
	}

	if _, err = r.CreateTag(article.TagList, res); err != nil {
		return nil, err
	}

	res.TagList = article.TagList
	res.Author = author

	return
}

func (r *articleRepo) CreateTag(tag_names []string, article *biz.ArticleReply) (tags []*biz.Tag, err error) {
	if err = r.data.db.Where("article_id = ?", article.ID).Delete(&biz.ArticleTag{}).Error; err != nil {
		return nil, err
	}

	for _, t := range tag_names {
		var tag *biz.Tag
		if err = r.data.db.Where(&biz.Tag{
			Name: t,
		}).FirstOrCreate(&tag).Error; err != nil {
			return nil, err
		}

		if err = r.data.db.Create(&biz.ArticleTag{
			ArticleId: int64(article.ID),
			TagId:     tag.Id,
		}).Error; err != nil {
			return nil, err
		}
	}

	return
}

func (r *articleRepo) UpdateArticle(article *biz.CreateReq, slug string) (res *biz.ArticleReply, err error) {
	var author *biz.User
	if err = r.data.db.Where("slug = ?", slug).Updates(&biz.Article{
		Title:       article.Title,
		Description: article.Description,
		Body:        article.Body,
	}).Find(&res).Error; err != nil {
		return nil, err
	}

	if err = r.data.db.Table("users").Where("id = ?", rwmiddleware.UserClaims.ID).Find(&author).Error; err != nil {
		return nil, err
	}

	if _, err = r.CreateTag(article.TagList, res); err != nil {
		return nil, err
	}

	res.TagList = article.TagList
	res.Author = author

	return
}

func (r *articleRepo) DeleteArticle(slug string) (err error) {
	var article *biz.Article

	if err = r.data.db.Where("slug = ?", slug).Delete(&article).Error; err != nil {
		return err
	}

	// if err = r.data.db.Where("slug = ?", slug).Scan(&article).Error; err != nil {
	// 	return err
	// }

	// if err = r.data.db.Where("article_id = ?", article.ID).Delete(&biz.ArticleTag{}).Error; err != nil {
	// 	return err
	// }

	// if err = r.data.db.Where("article_id = ?", article.ID).Delete(&biz.ArticleFavorite{}).Error; err != nil {
	// 	return err
	// }

	return
}

func (r *articleRepo) FavoriteArticle(slug string) (res *biz.ArticleReply, err error) {
	var author *biz.User

	if err = r.data.db.Table("articles").Where("slug = ?", slug).Find(&res).Error; err != nil {
		return nil, err
	}

	if err = r.data.db.Table("users").Where("id = ?", rwmiddleware.UserClaims.ID).Find(&author).Error; err != nil {
		return nil, err
	}

	if err = r.data.db.Where(&biz.ArticleFavorite{
		ArticleId: int64(res.ID),
		UserId:    author.Id,
	}).FirstOrCreate(&biz.ArticleFavorite{}).Error; err != nil {
		return nil, err
	}

	if err = r.WithTagsFavoritesCount([]*biz.ArticleReply{res}, author); err != nil {
		return nil, err
	}

	return
}

func (r *articleRepo) UnfavoriteArticle(slug string) (res *biz.ArticleReply, err error) {
	var author *biz.User
	if err = r.data.db.Table("articles").Where("slug = ?", slug).Find(&res).Error; err != nil {
		return nil, err
	}

	if err = r.data.db.Where("article_id = ?", res.ID).Delete(&biz.ArticleFavorite{}).Error; err != nil {
		return nil, err
	}

	if err = r.data.db.Table("users").Where("id = ?", rwmiddleware.UserClaims.ID).Find(&author).Error; err != nil {
		return nil, err
	}

	if err = r.WithTagsFavoritesCount([]*biz.ArticleReply{res}, author); err != nil {
		return nil, err
	}

	return
}

func (r *articleRepo) FindItem(items []int64, item int64) (index int, ok bool) {
	for i, t := range items {
		if t == item {
			return i, true
		}
	}

	return -1, false
}
