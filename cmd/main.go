package main

import (
	"github.com/gin-gonic/gin"
	"go-admin/internal/interfaces"
	"log"
)

func main() {
	log.Printf("Starting....")
	engine := gin.Default()
	interfaces.InitRouters(engine)
	engine.Run(":3200")
}
