package controllers

import (
	"tubes/apimodels"
	"tubes/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Registrasi(ctx *gin.Context) {
	//uid := ""
	var req apimodels.ReqRegistration
	res := apimodels.ResRegistration{
		Status:  -1,
		Message: "Failed",
	}
	uid, _ := ctx.Get("uuid")
	log.Printf("ID : [%v] Controller-Registrasi ", uid)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, res)
		ctx.Abort()
	}
	//
	//err := validate.Struct(user)
	//if err != nil {

	ctx.JSON(http.StatusAccepted, services.Registrasi(req))
}

func Login(ctx *gin.Context) {
	var req apimodels.ReqLogin
	res := apimodels.ResLogin{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, res)
		ctx.Abort()
	}

	token, err := services.GetToken(req)
	if err != nil {
		ctx.JSON(http.StatusFailedDependency, res)
		ctx.Abort()
	}
	res.Token = token
	ctx.JSON(http.StatusOK, res)
}
