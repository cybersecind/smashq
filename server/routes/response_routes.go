package routes

import (
    "github.com/gin-gonic/gin"
    "smashq/controllers"
)

func ResponseRoutes(router *gin.Engine) {
    responseGroup := router.Group("/response")
    {
        responseGroup.GET("/:agentID", controllers.ResponseHandler)

    }
}