package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os/exec"
	"restfulsdk/network"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"content-type"},
	}))
	fabricRouter := router.Group("/fabric")
	{
		channelRouter := fabricRouter.Group("/channels")
		{
			channelRouter.GET("/")
			channelRouter.POST("/create", network.CreateChannel)
			channelRouter.POST("/join", network.JoinChannel)
		}
		chaincodeRouter := fabricRouter.Group("/chaincodes")
		{
			chaincodeRouter.POST("/package", network.PackageCC)
			chaincodeRouter.GET("/install", network.InstalledCC)
			chaincodeRouter.POST("/install", network.InstallCC)
			chaincodeRouter.GET("/approve", network.ApprovedCC)
			chaincodeRouter.POST("/approve", network.ApproveCC)
			chaincodeRouter.GET("/commit", network.CommittedCC)
			chaincodeRouter.POST("/commit", network.CommitCC)
			chaincodeRouter.POST("/init", network.InitCC)
			chaincodeRouter.POST("/test", network.TestCC)
		}

	}

	return router
}

func main() {

	// init sdk env info

	updateSwagger()
	r := setupRouter()
	r.Run(":8080")
}
func setUp() {

}
func updateSwagger() {
	cmd := exec.Command("swag", "init", "-g", "main.go", "--output", "docs")
	cmd.Run()
}
