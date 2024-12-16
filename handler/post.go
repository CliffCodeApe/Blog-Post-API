package handler

import (
	"blog_post/contract"
	"blog_post/dto"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	service contract.PostService
}

func NewPostHandler(service contract.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func extractUser(c *gin.Context) (uint64, string, error) {
	userID, exists := c.Get("userId")
	if !exists {
		return 0, "", errors.New("user ID not found in context")
	}

	username, exists := c.Get("username")
	if !exists {
		return 0, "", errors.New("username not found in context")
	}

	return userID.(uint64), username.(string), nil
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var postDTO dto.PostRequest
	if err := c.ShouldBindJSON(&postDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract user ID and username from the token
	userID, username, err := extractUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Create the post using the extracted user information
	if err := h.service.CreatePost(postDTO, userID, username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post created successfully"})
}

func (h *PostHandler) GetPosts(c *gin.Context) {
	posts, _ := h.service.GetPosts()
	c.JSON(http.StatusOK, posts)
}
