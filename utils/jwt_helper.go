package utils

import (
	"crypto/rsa"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID   uint64 `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

// Load the private key from the specified file
func LoadPrivateKey(filePath string) (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile(filePath) // Use os.ReadFile instead of ioutil
	if err != nil {
		return nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// Load the public key from the specified file
func LoadPublicKey(filePath string) (*rsa.PublicKey, error) {
	keyData, err := os.ReadFile(filePath) // Use os.ReadFile instead of ioutil
	if err != nil {
		return nil, err
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

// Initialize the keys (call this function at the start of your application)
func InitKeys() error {
	privateKeyPath := os.Getenv("PRIVATE_KEY") // Get the private key path from the environment variable
	publicKeyPath := os.Getenv("PUBLIC_KEY")   // Get the public key path from the environment variable

	var err error
	privateKey, err = LoadPrivateKey(privateKeyPath)
	if err != nil {
		return err
	}

	publicKey, err = LoadPublicKey(publicKeyPath)
	if err != nil {
		return err
	}

	return nil
}

func GenerateJWT(userID uint64, username string, email string) (string, error) {
	claims := jwt.MapClaims{
		"userId":   userID,
		"username": username,
		"email":    email,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // Token expiration
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Ensure privateKey is not nil
	if privateKey == nil {
		return "", errors.New("private key is not initialized")
	}

	return token.SignedString(privateKey)
}

func DecodeJWT(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil // Use the public key for verification
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
