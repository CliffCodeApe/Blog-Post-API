package handler

import (
	"blog_post/contract"
	"blog_post/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postService contract.PostService
}

func NewPostHandler(service contract.PostService) *PostHandler {
	return &PostHandler{postService: service}
}

// CreatePost godoc
// @Summary Creates Post
// @Description Creates a new post
// @Tags posts
// @Accept json
// @Produce json
// @Param post body dto.PostRequest true "Post data"
// @Success 201 {object} dto.PostResponse "Post Created Successfully"
// @Failure 400 {object} dto.InvalidInputErrorResponse "Invalid input data"
// @Failure 404 {object} dto.NotFoundErrorResponse "Resource not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error
// @Router /posts/ [post]
func (h *PostHandler) CreatePost(c *gin.Context) {
	var postDTO dto.PostRequest
	if err := c.ShouldBindJSON(&postDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.postService.CreatePost(postDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "201", "message": "Post created successfully"})
}

// GetAllPosts godoc
// @Summary Get All Posts
// @Description Get all the posts
// @Tags posts
// @Accept json
// @Produce json
// @Success 200 {object} dto.GetResponse "Post retrieved successfully"
// @Failure 404 {object} dto.NotFoundErrorResponse "Post not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /posts/ [get]
func (h *PostHandler) GetAllPosts(c *gin.Context) {
	posts, _ := h.postService.GetAllPosts()
	c.JSON(http.StatusOK, posts)
}

// GetPostByID godoc
// @Summary Get Post by ID
// @Description Retrieve a post by its ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} dto.GetResponse "Post retrieved successfully"
// @Failure 404 {object} dto.NotFoundErrorResponse "Post not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /posts/{id} [get]
func (h *PostHandler) GetPostByID(c *gin.Context) {
	postId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	post, err := h.postService.GetPostByID(uint64(postId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// UpdatePost godoc
// @Summary Update Post
// @Description Update an existing post
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Param post body dto.EditRequest true "Post data"
// @Success 200 {object} dto.EditResponse "Post updated successfully"
// @Failure 400 {object} dto.InvalidInputErrorResponse "Invalid input data"
// @Failure 404 {object} dto.NotFoundErrorResponse "Post not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /posts/{id} [put]
// @Router /posts/{id} [patch]
func (h *PostHandler) UpdatePost(c *gin.Context) {
	var req dto.EditRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	if err := h.postService.UpdatePost(uint64(postId), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "200", "message": "Post updated successfully"})
}

// DeletePost godoc
// @Summary Delete Post
// @Description Delete a post by its ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} dto.DeleteResponse "Post updated successfully"
// @Failure 404 {object} dto.NotFoundErrorResponse "Post not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /posts/{id} [delete]
func (h *PostHandler) DeletePost(c *gin.Context) {
	postId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	if err := h.postService.DeletePost(uint64(postId)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "204", "message": "Post deleted successfully"})
}
