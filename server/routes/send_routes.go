package routes

import (
    "github.com/gin-gonic/gin"
    "smashq/controllers"
)

func SendRoutes(router *gin.Engine) {
    sendGroup := router.Group("/send")
    {
        sendGroup.POST("/", controllers.SendHandler)

    }
}