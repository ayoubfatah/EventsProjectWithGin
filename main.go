package main

import (
	"gin-quickstart/db"
	"gin-quickstart/routes"

	"github.com/gin-gonic/gin"
)


func main(){
	 db.InitDB()

 	server :=	gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}

