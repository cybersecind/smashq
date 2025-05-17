package main

import (
    "fmt"
    "time"
    "github.com/gin-gonic/gin"
    "smashq/mqtt"
    "smashq/routes"
    "github.com/gin-contrib/cors"
)

func main() {
    mqtt.Init()

    router := gin.Default()

    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*", "http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    routes.SendRoutes(router)
    routes.ResponseRoutes(router)

    fmt.Println("C2 server running at http://localhost:8080")
    router.Run(":8080")
}
