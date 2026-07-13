package repository

import (
	"context"

	"github.com/awanda/backend-repo/internal/model"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(ctx context.Context, post *model.Post) error
	FindAll(ctx context.Context, limit, offset int) ([]model.Post, error)
	FindByID(ctx context.Context, id int) (*model.Post, error)
	Update(ctx context.Context, post *model.Post) error
	Delete(ctx context.Context, id int) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Create(ctx context.Context, post *model.Post) error {
	return r.db.WithContext(ctx).Omit("CreatedDate", "UpdatedDate").Create(post).Error
}

func (r *postRepository) FindAll(ctx context.Context, limit, offset int) ([]model.Post, error) {
	var posts []model.Post
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&posts).Error
	return posts, err
}

func (r *postRepository) FindByID(ctx context.Context, id int) (*model.Post, error) {
	var post model.Post
	err := r.db.WithContext(ctx).First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) Update(ctx context.Context, post *model.Post) error {
	return r.db.WithContext(ctx).Model(post).Updates(map[string]interface{}{
		"title":    post.Title,
		"content":  post.Content,
		"category": post.Category,
		"status":   post.Status,
	}).Error
}

func (r *postRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&model.Post{}, id).Error
}
