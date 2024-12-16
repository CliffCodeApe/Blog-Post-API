package main

import (
	"blog_post/handler"
	"blog_post/repository"
	"blog_post/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

}

func main() {

	// Get configuration from environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Initialize database connection
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize repositories
	postRepo := repository.NewPostRepository(db)

	// Initialize services
	postService := service.NewPostService(postRepo)

	// Initialize handlers
	postHandler := handler.NewPostHandler(postService)

	// Setup routes
	r := gin.Default()

	protected := r.Group("/posts")
	protected.GET("/", postHandler.GetPosts)
	protected.POST("/", postHandler.CreatePost)

	// Start the server
	log.Printf("Starting server on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
