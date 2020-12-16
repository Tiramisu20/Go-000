package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"project/internal/service"
)

func main() {
	log.Printf("start!\n")

	err := Start()

	if err != nil {
		log.Printf("startup failed: %v\n", err)
	}
}

// TODO resource init
// * db init
// * cache init

// TODO app start

// TODO Learn gRPC

// TODO Learn Wire


func Start() error {
	r := gin.Default()
	registerRouter(r)
	return r.Run()
}

func registerRouter(r *gin.Engine) {

	r.GET("/listUsers", func(c *gin.Context) {
		users := service.NewUserService().ListUsers()
		result, _ := json.Marshal(users)
		c.JSON(200, gin.H{
			"data" : result,
		})
	})
}



