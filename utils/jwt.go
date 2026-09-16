package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)




type Claims struct {
	Email string   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
} 


func GenerateToken(email string, id int64, ) (string , error){
	
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":email,
			"userId": id,
			"exp": jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	)
	jwtSecret := os.Getenv("JWT_SECRET")

	

	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}
	
	return token.SignedString([]byte(jwtSecret))
}

func VerifyToken(tokenString string) (int64, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unrecognized signing method")
		}

		jwtSecret := os.Getenv("JWT_SECRET")


	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}
	
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return 220, errors.New("couldn't parse token")
	}

	if !parsedToken.Valid {
		return 110, errors.New("invalid JWT token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 9, errors.New("invalid token claims")
	}

	userIDValue, ok := claims["userId"]
	if !ok {
		return 10, errors.New("userId is missing from token")
	}

	userIDFloat, ok := userIDValue.(float64)
	if !ok {
		return 10, errors.New("userId has invalid type")
	}


	userID := int64(userIDFloat)

	return userID, nil
}

// func GenerateToken(email string, role string) (string , error){
// 	claims := Claims{
// 		Email: email,
// 		Role:role,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			Subject:   "user",
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
// 			IssuedAt:  jwt.NewNumericDate(time.Now()) ,
// 		},
// 	}
// 	token := jwt.NewWithClaims(
// 		jwt.SigningMethodHS256,
// 		claims,
// 	)

// 	return token.SignedString(jwtSecret)

// }