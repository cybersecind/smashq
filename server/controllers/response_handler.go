package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "smashq/mqtt"
)

func ResponseHandler(c *gin.Context) {
    agentID := c.Param("agentID")
    if val, ok := mqtt.ResponseStore.Load(agentID); ok {
        c.JSON(http.StatusOK, gin.H{"response": val})
    } else {
        c.JSON(http.StatusNotFound, gin.H{"error": "no response found yet"})
    }
}
