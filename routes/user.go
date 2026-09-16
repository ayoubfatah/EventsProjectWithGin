package routes

import (
	"gin-quickstart/models"
	"gin-quickstart/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)


func createUser(c *gin.Context){
	var user models.User

err :=c.ShouldBindJSON(&user)
 if err!=nil{
		 c.JSON(http.StatusBadRequest, gin.H{
            "error": "something happened while trying to parse the data ",
        })
        return
}
 	err =  user.Save()
 	if(err !=nil){
		c.JSON(http.StatusInternalServerError , gin.H{
		"message": err.Error(),
		})
		return 
	}
  
  c.JSON(http.StatusCreated, gin.H{
	"message":"User Created",
  })
} 

func getAllUsers(c *gin.Context){
	users , err :=	models.GetUsers()
	  
	if(err !=nil){
		c.JSON(http.StatusInternalServerError , gin.H{
			"message": "an error happened while trying to get users",
		
		})
		return
	}
	  c.JSON(http.StatusOK, gin.H{ 
      "users": users,
    }) 
}

func login(c *gin.Context){
	var user models.SignupRequest 
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "invalid credentials", 
		})
		return
	}



	token, err := utils.GenerateToken(user.Email, user.Id)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "could not generate token",
		})
		return
	}


	if err != nil {
		c.JSON(500, gin.H{
			"error": "could not generate token",
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
		"user":user,
	})
   
}




func getCurrentUser(c *gin.Context){

    userId := c.GetInt64("userId")



    user, err := models.GetUserByID(userId)

	


    if err != nil {
        c.JSON(404, gin.H{
            "error": "User not found",
        })
        return
    }

    c.JSON(200, user)

}