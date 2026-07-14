package service

import (
	"context"
	"errors"

	"github.com/awanda/backend-repo/internal/dto"
	"github.com/awanda/backend-repo/internal/model"
	"github.com/awanda/backend-repo/internal/repository"
	"github.com/awanda/backend-repo/internal/validator"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("article not found")

type PostService interface {
	Create(ctx context.Context, req dto.PostRequest) error
	GetList(ctx context.Context, limit, offset int) ([]dto.PostResponse, error)
	GetDetail(ctx context.Context, id int) (*dto.PostResponse, error)
	Update(ctx context.Context, id int, req dto.PostRequest) error
	Delete(ctx context.Context, id int) error
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{repo: repo}
}

func (s *postService) Create(ctx context.Context, req dto.PostRequest) error {
	if err := validator.ValidatePostRequest(req); err != nil {
		return err
	}

	post := &model.Post{
		Title:    req.Title,
		Content:  req.Content,
		Category: req.Category,
		Status:   req.Status,
	}

	return s.repo.Create(ctx, post)
}

func (s *postService) GetList(ctx context.Context, limit, offset int) ([]dto.PostResponse, error) {
	posts, err := s.repo.FindAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.PostResponse, 0, len(posts))
	for _, post := range posts {
		responses = append(responses, toPostResponse(post))
	}

	return responses, nil
}

func (s *postService) GetDetail(ctx context.Context, id int) (*dto.PostResponse, error) {
	post, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	response := toPostResponse(*post)
	return &response, nil
}

func (s *postService) Update(ctx context.Context, id int, req dto.PostRequest) error {
	if err := validator.ValidatePostRequest(req); err != nil {
		return err
	}

	post, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound
		}
		return err
	}

	post.Title = req.Title
	post.Content = req.Content
	post.Category = req.Category
	post.Status = req.Status

	return s.repo.Update(ctx, post)
}

func (s *postService) Delete(ctx context.Context, id int) error {
	post, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound
		}
		return err
	}

	return s.repo.Delete(ctx, post.ID)
}

func toPostResponse(post model.Post) dto.PostResponse {
	return dto.PostResponse{
		ID:       post.ID,
		Title:    post.Title,
		Content:  post.Content,
		Category: post.Category,
		Status:   post.Status,
	}
}
