package validator

import (
	"errors"

	"github.com/awanda/backend-repo/internal/dto"
)

var ErrValidation = errors.New("validation failed")

func ValidatePostRequest(req dto.PostRequest) error {
	if req.Title == "" || len(req.Title) < 20 {
		return ErrValidation
	}
	if req.Content == "" || len(req.Content) < 200 {
		return ErrValidation
	}
	if req.Category == "" || len(req.Category) < 3 {
		return ErrValidation
	}
	if req.Status != "publish" && req.Status != "draft" && req.Status != "thrash" {
		return ErrValidation
	}
	return nil
}
