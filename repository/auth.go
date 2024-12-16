package repository

import (
	"blog_post/entity"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) SaveUser(user entity.User) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}

	// Log the saved user for debugging
	log.Printf("User registered: %+v", user)
	return nil
}

func (r *AuthRepository) GetUserByEmail(email string, password string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err // User not found
	}

	// Log the user found for debugging
	log.Printf("User found: %+v", user)

	// Compare the provided password with the hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("invalid credentials") // Password does not match
	}

	return user, nil
}
