package main

import (
	bucketController "api/controllers/bucket"
	instanceController "api/controllers/instance"
	networkController "api/controllers/network"
	"api/seed"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	seed.Seed()

	server := gin.New()

	server.Use(gin.Recovery())
	// config := cors.DefaultConfig()

	server.Use(cors.Default())

	server.GET("/postman", func(ctx *gin.Context) {
		jsonFile, err := os.Open("./postman/postman_collection.json")
		if err != nil {
			log.Err(err)
		}
		log.Print("Successfully Opened postman_collection.json")

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var postman interface{}
		json.Unmarshal(byteValue, &postman)

		ctx.JSON(http.StatusOK, postman)

	})

	server.GET("/swagger", func(ctx *gin.Context) {
		jsonFile, err := os.Open("./postman/swagger.json")
		if err != nil {
			log.Err(err)
		}
		log.Print("Successfully Opened swagger.json")

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var swagger interface{}
		json.Unmarshal(byteValue, &swagger)
		ctx.JSON(http.StatusOK, swagger)

	})

	server.POST("/signin", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"token": "signin",
		})
	})

	server.POST("/instance/read", instanceController.Read)
	server.POST("/instance", instanceController.Create)
	server.PUT("/instance", instanceController.Update)
	server.DELETE("/instance", instanceController.Delete)

	server.POST("/bucket/read", bucketController.Read)
	server.POST("/bucket", bucketController.Create)
	server.PUT("/bucket", bucketController.Update)
	server.DELETE("/bucket", bucketController.Delete)

	server.POST("/network/read", networkController.Read)
	server.POST("/network", networkController.Create)
	server.PUT("/network", networkController.Update)
	server.DELETE("/network", networkController.Delete)

	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:8090/swagger")))

	// server.GET("/ws", func(c *gin.Context) {
	// 	wshandler(c.Writer, c.Request)
	// })

	server.Run(":8090")
}
