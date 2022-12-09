package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os/exec"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"content-type"},
	}))
	router.POST("channel", dashboard.AddDataToES)
	return router
}

func main() {
	updateSwagger()
	r := setupRouter()
	r.Run(":8080")
}
func updateSwagger() {
	cmd := exec.Command("swag", "init", "-g", "main.go", "--output", "docs")
	cmd.Run()
}
