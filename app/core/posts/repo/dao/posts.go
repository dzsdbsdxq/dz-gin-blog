package dao

import (
	"errors"
	"fmt"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"gorm.io/gorm"
)

type IPostDao interface {
	Create(post *model.SysPosts) error
	GetPostDetailBySlug(slug string) (sysPost *model.SysPosts, err error)
	GetPostDetailById(id uint) (sysPost *model.SysPosts, err error)
	Update(post *model.SysPosts) error
	DeletePostByIds(ids global.IdsReq) error
}

var _ IPostDao = (*PostDao)(nil)

func NewPostDao() *PostDao {
	return &PostDao{
		coll: global.G_DZ_DB,
	}
}

type PostDao struct {
	coll *gorm.DB
}

//@author: [jackieLee](https://github.com/dzsdbsdxq)
//@function: Create
//@description: 创建文章
//@param: post *model.SysPosts
//@return: error

func (p *PostDao) Create(post *model.SysPosts) error {
	_, err := p.GetPostDetailBySlug(post.Slug)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("存在相同的别名:%s", post.Slug))
	}
	return p.coll.Create(&post).Error
}

//@author: [jackieLee](https://github.com/dzsdbsdxq)
//@function: GetPostDetailBySlug
//@description: 根据别名slug获取文章
//@param: slug string
//@return: sysPost model.SysPosts, err error

func (p *PostDao) GetPostDetailBySlug(slug string) (sysPost *model.SysPosts, err error) {
	err = p.coll.Where("slug = ?", slug).First(&sysPost).Error
	return
}

//@author: [jackieLee](https://github.com/dzsdbsdxq)
//@function: GetPostDetailById
//@description: 根据id获取文章
//@param: id uint
//@return: sysPost model.SysPosts, err error

func (p *PostDao) GetPostDetailById(id uint) (sysPost *model.SysPosts, err error) {
	err = p.coll.Where("id = ?", id).First(&sysPost).Error
	return
}

//@author: [jackieLee](https://github.com/dzsdbsdxq)
//@function: Update
//@description: 更新文章
//@param: id uint
//@return: sysPost model.SysPosts, err error

func (p *PostDao) Update(post *model.SysPosts) error {
	_, err := p.GetPostDetailBySlug(post.Slug)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("存在相同的别名:%s", post.Slug))
	}
	return p.coll.Save(&post).Error
}

//@author: [jackieLee](https://github.com/dzsdbsdxq)
//@function: DeletePostByIds
//@description: 批量删除文章，加入回收站
//@param: ids vo.IdsReq
//@return: error

func (p *PostDao) DeletePostByIds(ids global.IdsReq) error {
	var posts []model.SysPosts
	return p.coll.Find(&posts, "id in ?", ids.Ids).Delete(&posts).Error
}

func (p *PostDao) GetPostsList(postReq vo.PostReq, info global.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	//limit := info.PageSize
	//offset := info.PageSize * (info.PageNo - 1)
	//db := p.coll.Model(&model.SysPosts{})
	//var postList []model.SysPosts
	//if info.Keyword != "" {
	//	db = db.Where("title LIKE ?", "%"+info.Keyword+"%")
	//}
	//if info.Category != 0 {
	//
	//}
	//
	//err = db.Count(&total).Error
	return nil, 0, nil

}
