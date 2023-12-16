package util

import (
	"e-wallet/dto"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"os"
	"time"
)

func GenerateRandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	var charsets = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890")
	letters := make([]rune, n)
	for i := range letters {
		letters[i] = charsets[rand.Intn(len(charsets))]
	}
	return string(letters)
}

func GenerateRandomNumber(n int) string {
	rand.Seed(time.Now().UnixNano())
	var charsets = []rune("01234567890")
	letters := make([]rune, n)
	for i := range letters {
		letters[i] = charsets[rand.Intn(len(charsets))]
	}
	return string(letters)
}

var JwtKey = []byte(os.Getenv("JWT_KEY"))

func GenerateJWTToken(user dto.UserData) (string, error) {

	// Membuat struktur klaim JWT
	accessTokenExpiration := time.Now().Add(10 * time.Minute)
	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   accessTokenExpiration,
	}

	// Membuat token JWT dengan klaim dan tanda tangan
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
