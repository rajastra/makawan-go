package services

import (
	"tubes/apimodels"
	"tubes/database"
	"tubes/models"
)

func SaveProduct(req apimodels.RequestProduct) (apimodels.ResponseProduct, error) {
	var res apimodels.ResponseProduct
	PrettyPrint(req)
	db := database.GetDb()
	
	product := models.Product{
		ProductName: req.ProductName,
		Price:       req.Price,
		Description: req.Description,
	}

	errdb := db.Create(&product).Error

	if errdb != nil {
		return res, errdb
	}

	return apimodels.ResponseProduct{
		Data:         req,
		ProductID:    product.ProductID,
		ResponseCode: "00",
		Status:       "Success",
		Total:        product.Price,
	}, nil
	
}

func GetAll() (apimodels.ResponseGeneral, error) {
	var res apimodels.ResponseGeneral
	db := database.GetDb()
	var products []models.Product
	errdb := db.Find(&products).Error
	if errdb != nil {
		return res, errdb
	}
	return apimodels.ResponseGeneral{
		Meta: apimodels.Meta{
			Code:    0,
			Message: "Success",
			Status:  "Success",
		},
		Data: products,
	}, nil
}