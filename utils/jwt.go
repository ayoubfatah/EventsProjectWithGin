package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("super-secret-key")


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
	return token.SignedString(jwtSecret)
	return  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNzc0BnbWFpbC5jb20iLCJleHAiOjE3ODg2ODYxNjQsInVzZXJJZCI6MX0.c8YCfo7_xggj8PVllE2QfReImjwO0KrxUPMO5OQv1Vw" ,nil
}



func VerifyToken(token string )(int64 , error){

	 parsedToken , err :=	jwt.Parse(token, func(token *jwt.Token)(any , error){
			_ , ok := token.Method.(*jwt.SigningMethodHMAC)

			if(!ok){
				return nil ,errors.New("Unrecognized method")
			}
			return jwtSecret , nil
		})
		if err !=nil {
				return 0,errors.New("Couldn't parse the token")

			}
	tokenIsValid := parsedToken.Valid		
	if !tokenIsValid {
				return 0,errors.New("Invalid jwt token")

	}
	claims  , ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
				return  0,errors.New("Invalid token claims ")
	}
	// email :=  claims["email"].(string)
	userId := int64(claims["userId"].(float64))




	return   userId,nil 
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