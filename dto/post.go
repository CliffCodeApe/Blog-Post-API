package dto

import "time"

type GetPost struct {
	ID        int       `json:"post_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type EditRequest struct {
	Title   *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
}

type GetResponse struct {
	Status  int    `json:"status" example:"200"`
	Message string `json:"Message" example:"Post Successfully Retreived"`
}

type PostResponse struct {
	Status  int    `json:"status" example:"201"`
	Message string `json:"Message" example:"Post Successfully Created"`
}

type EditResponse struct {
	Status  int    `json:"status" example:"200"`
	Message string `json:"Message" example:"Post Successfully Updated"`
}

type DeleteResponse struct {
	Status  int    `json:"status" example:"204"`
	Message string `json:"Message" example:"Post Successfully Deleted"`
}

// Error Respones

// ErrorResponse represents a generic error response
type ErrorResponse struct {
	Error string `json:"error" example:"An error occurred"`
}

// InvalidInputErrorResponse represents an error for invalid input
type InvalidInputErrorResponse struct {
	ErrorCode int    `json:"error_code" example:"400"`
	Message   string `json:"message" example:"Invalid input data"`
}

// NotFoundErrorResponse represents an error for not found resources
type NotFoundErrorResponse struct {
	ErrorCode int    `json:"error_code" example:"404"`
	Message   string `json:"message" example:"Resource not found"`
}

// InternalServerErrorResponse represents an internal server error
type InternalServerErrorResponse struct {
	ErrorCode int    `json:"error_code" example:"500"`
	Message   string `json:"message" example:"Internal server error"`
}
