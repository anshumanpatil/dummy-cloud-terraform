package main

import (
	authController "api/controllers/auth"
	bucketController "api/controllers/bucket"
	instanceController "api/controllers/instance"
	networkController "api/controllers/network"
	"api/jwthelper"
	"api/kf2wbsock"
	"api/seed"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	seed.Seed()

	go startWebKafka()
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
		ctx := context.Background()

		reqUrlPath := c.Request.URL.String()

		if !(strings.Contains(reqUrlPath, "read")) {
			produce(ctx, "change", c.Request.URL.String())
			fmt.Println("change", c.Request.URL.String())
		}

		c.Next()
	}
}

func startWebKafka() {
	configFile := flag.String("config", "config.yaml", "Config file location")
	// kf2wbsock
	list := kf2wbsock.ReadK2WS(*configFile)

	for i := range list {
		go func(k2ws *kf2wbsock.K2WS) {
			err := k2ws.Start()
			if err != nil {
				// log.Fatalln(err)
			}
		}(list[i])
	}

	var chExit = make(chan os.Signal, 1)
	signal.Notify(chExit, os.Interrupt)
	<-chExit
}

const (
	topic          = "anshu"
	broker1Address = "localhost:9092"
	// broker2Address = "localhost:9094"
	// broker3Address = "localhost:9095"
)

func produce(ctx context.Context, title string, msg string) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Address},
		Topic:   topic,
	})

	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(title),
		// create an arbitrary message payload for the value
		Value: []byte(msg),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}
}
