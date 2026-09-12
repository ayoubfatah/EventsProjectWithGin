package main

import (
	"gin-quickstart/db"
	"gin-quickstart/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main(){
	db.InitDB()
	server :=	gin.Default()

	server.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:3000"},
    AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
    AllowCredentials: true,
}))
	// models.SeedEvents()
	



	routes.RegisterRoutes(server)
	server.Run(":8080")
}

