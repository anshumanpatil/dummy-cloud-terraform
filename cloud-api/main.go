package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.Use(dbFlushMiddleware())
	r.POST("/api/login", loginHandler)
	r.POST("/api/secret/register", registerHandler)
	r.GET("/api/users", usersHandler)
	r.GET("/api/ping", pingHandler)

	r.GET("/api/infrastructures", getAllInfrastructuresHandler)

	r.POST("/api/infrastructure/single", getInfrastructuresHandler)
	r.DELETE("/api/infrastructure/single", deleteInfrastructuresHandler)

	r.POST("/api/infrastructure", setInfrastructuresHandler)
	r.PUT("/api/infrastructure/single", updateInfrastructuresHandler)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080

}
