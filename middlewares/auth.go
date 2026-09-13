package middlewares

import (
	"fmt"
	"gin-quickstart/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticated(c *gin.Context){
	token := c.Request.Header.Get("Authorization")
 fmt.Println("token")
 fmt.Println(token)

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header required",
			})
			return
		}
	
	userId,err :=	utils.VerifyToken(token)
	
	fmt.Println("user id from auth")
	fmt.Println(userId)
		fmt.Println("error 1 ")
		fmt.Println(err)



	if err !=nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not authorized",

			}) 	
		return
		}

		
		fmt.Println("error 2 ")
		fmt.Println(err)	

	c.Set("userId", userId)
	value, exists := c.Get("userId")

fmt.Println("EXISTS:", exists)
fmt.Printf("VALUE: %v\n", value)
fmt.Printf("TYPE: %T\n", value)
	c.Next()

}