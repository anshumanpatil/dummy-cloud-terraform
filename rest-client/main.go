package main

import (
	authController "api/controllers/auth"
	bucketController "api/controllers/bucket"
	instanceController "api/controllers/instance"
	networkController "api/controllers/network"
	"api/jwthelper"
	"api/seed"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

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

	server.Use(CORSMiddleware())
	// publicEndpoints := r.Group("/group1", middleware1())
	privateEndpoints := server.Group("/", jwthelper.AuthMiddleware())

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

	server.POST("/signin", authController.Login)

	privateEndpoints.POST("/instance/read", instanceController.Read)
	privateEndpoints.POST("/instance", instanceController.Create)
	privateEndpoints.PUT("/instance", instanceController.Update)
	privateEndpoints.DELETE("/instance", instanceController.Delete)

	privateEndpoints.POST("/bucket/read", bucketController.Read)
	privateEndpoints.POST("/bucket", bucketController.Create)
	privateEndpoints.PUT("/bucket", bucketController.Update)
	privateEndpoints.DELETE("/bucket", bucketController.Delete)

	privateEndpoints.POST("/network/read", networkController.Read)
	privateEndpoints.POST("/network", networkController.Create)
	privateEndpoints.PUT("/network", networkController.Update)
	privateEndpoints.DELETE("/network", networkController.Delete)

	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:8090/swagger")))

	// server.GET("/ws", func(c *gin.Context) {
	// 	wshandler(c.Writer, c.Request)
	// })

	server.Run(":8090")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
