package controllers

import (
	"tubes/apimodels"
	"tubes/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	var req apimodels.Request
	var res apimodels.Response
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := services.SaveOrder(req)
	if err != nil {
		res.Status = "Create Order Gagal"
		res.ResponseCode = "400"
	}
	ctx.JSON(http.StatusOK, res)
	return
}

// this
func GetOrderBy(ctx *gin.Context) {
	orderId := ctx.Param("orderID")
	log.Println("OrderID :", orderId)
	res, _ := services.GetOrderBy(orderId)
	ctx.JSON(http.StatusOK, res)
}

func UpdateOrder(ctx *gin.Context) {
	orderId := ctx.Param("orderID")
	log.Println("OrderID :", orderId)
	var req apimodels.Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	log.Println("Request :", req)
	res, err := services.UpdateOrder(req, orderId)
	if err != nil {
		res.Status = "Update Order Gagal"
		res.ResponseCode = "400"
	}
	ctx.JSON(http.StatusOK, res)
}

func DeleteOrder(ctx *gin.Context) {
	orderId := ctx.Param("orderID")
	log.Println("OrderID :", orderId)
	if orderId == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if services.DeleteOrder(orderId) {
		ctx.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		//ctx.Writer.WriteHeader(http.StatusNoContent)
		//ctx.Done()
		ctx.Status(http.StatusNoContent)
		return
	}
	ctx.AbortWithStatus(http.StatusExpectationFailed)

}
