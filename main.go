package main

import (
	"github.com/furkanerdogan/Go-Postgresql-Dockerized-Crud-Api/config"
	"github.com/furkanerdogan/Go-Postgresql-Dockerized-Crud-Api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":4545")
}
