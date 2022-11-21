package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func HealthCheck(ctx *gin.Context) {
	response := Response{
		Message: "UP",
	}
	ctx.JSON(http.StatusOK, response)
}
