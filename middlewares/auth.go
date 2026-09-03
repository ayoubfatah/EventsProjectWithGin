package middlewares

import (
	"fmt"
	"gin-quickstart/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticated(c *gin.Context){
	token := c.Request.Header.Get("Authorization")
		fmt.Printf("value n token %v",token)
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header required",
			})
			return
		}
	
	userId,err :=	utils.VerifyToken(token)


	if err !=nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not authorized",
			}) }

	c.Set("userId", userId)
	c.Next()

}