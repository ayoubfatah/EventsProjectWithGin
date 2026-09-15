package middlewares

import (
	"gin-quickstart/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticated(c *gin.Context){
	token := c.Request.Header.Get("Authorization")



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

			}) 	
		return
		}

		



	c.Set("userId", userId)
	c.Next()

}