package main

import (
	"client"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	usecase client.UseCase
)

func Verify(ctx *gin.Context) {
	body := new(client.TCKimlik)

	err := ctx.ShouldBindJSON(body)
	if err != nil {
		msg := fmt.Sprintf("Invalid request : %s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": msg,
		})
		return
	}

	verify := usecase.MakeRequest(body)
	ctx.JSON(http.StatusOK, gin.H{
		"isValid": verify,
	})
}

func Ping(ctx *gin.Context) {
	message := fmt.Sprintf("pong - %s", usecase.GenerateRandomString(5))

	ctx.JSON(http.StatusOK, gin.H{
		"meessage": message,
	})
}
