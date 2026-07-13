package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/awanda/backend-repo/internal/dto"
	"github.com/awanda/backend-repo/internal/service"
	"github.com/awanda/backend-repo/internal/validator"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	service service.PostService
}

func NewPostHandler(service service.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) Create(c *gin.Context) {
	var req dto.PostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "validation failed"})
		return
	}

	if err := h.service.Create(c.Request.Context(), req); err != nil {
		if errors.Is(err, validator.ErrValidation) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "validation failed"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *PostHandler) GetList(c *gin.Context) {
	limit, err := strconv.Atoi(c.Param("a"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "validation failed"})
		return
	}

	offset, err := strconv.Atoi(c.Param("b"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "validation failed"})
		return
	}

	posts, err := h.service.GetList(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	if posts == nil {
		posts = []dto.PostResponse{}
	}

	c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) GetDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("a"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "validation failed"})
		return
	}

	post, err := h.service.GetDetail(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "article not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("a"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "validation failed"})
		return
	}

	var req dto.PostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "validation failed"})
		return
	}

	if err := h.service.Update(c.Request.Context(), id, req); err != nil {
		if errors.Is(err, validator.ErrValidation) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "validation failed"})
			return
		}
		if errors.Is(err, service.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "article not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *PostHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("a"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "validation failed"})
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "article not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
