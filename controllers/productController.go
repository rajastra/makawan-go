package controllers

import (
	"net/http"
	"tubes/apimodels"
	"tubes/services"

	"github.com/gin-gonic/gin"
)

func CreateProduct(ctx *gin.Context) {
	var req apimodels.RequestProduct
	var res apimodels.ResponseProduct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := services.SaveProduct(req)
	if err != nil {
		res.Status = "Create Product Gagal"
		res.ResponseCode = "400"
	}
	ctx.JSON(http.StatusOK, res)
}

func GetAllProduct(ctx *gin.Context) {
	res, _ := services.GetAll()
	ctx.JSON(http.StatusOK, res)
}