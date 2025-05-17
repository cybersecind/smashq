package controllers

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
    "smashq/mqtt"
    MQTT "github.com/eclipse/paho.mqtt.golang"
)

func SendHandler(c *gin.Context) {
    var req struct {
        AgentID string `json:"agent_id"`
        Command string `json:"command"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    cmdTopic := fmt.Sprintf("agents/%s/cmd", req.AgentID)
    respTopic := fmt.Sprintf("agents/%s/resp", req.AgentID)

    // Subscribe once
    mqtt.Client.Subscribe(respTopic, 0, func(client MQTT.Client, msg MQTT.Message) {
        mqtt.ResponseStore.Store(req.AgentID, string(msg.Payload()))
    })

    mqtt.Client.Publish(cmdTopic, 0, false, req.Command)
    c.JSON(http.StatusOK, gin.H{"status": "command sent"})
}
