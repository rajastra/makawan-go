package services

import (
	"crypto/rand"
	"tubes/apimodels"
	"tubes/database"
	"tubes/models"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func SaveOrder(req apimodels.Request) (apimodels.Response, error) {
	var res apimodels.Response
	PrettyPrint(req)
	db := database.GetDb()
	var items []models.Item
	var total int64
	for _, vitem := range req.Items {
		var item models.Item
		item.Price = vitem.Price
		item.Quantity = int(vitem.Quantity)
		item.ItemCode = vitem.ItemCode
		item.Description = vitem.Description
		items = append(items, item)
		total += (vitem.Quantity * vitem.Price)
	}

	order := models.Order{
		OrderID:      generateRandomID(),
		CustomerName: req.CustomerName,
		OrderAt:      currentTime(),
		DetaiItem:    items,
	}

	errdb := db.Create(&order).Error
	if errdb != nil {
		return res, errdb
	}

	return apimodels.Response{
		Data:         req,
		DateTrans:    fmt.Sprintf("%v", dateTimeEpoch(currentTime())),
		OrderID:      order.OrderID,
		ResponseCode: "00",
		Status:       "Success",
		Total:        total,
	}, nil

}

func GetOrderBy(orderID string) (res apimodels.ResponseGeneral, err error) {
	res.Meta.Code = -1
	res.Meta.Message = "Failed"
	res.Meta.Status = "Failed"
	db := database.GetDb()
	order := &models.Order{}

	log.Println("Order: ", orderID)

	err = db.Preload("Items").Where("order_id = ?", orderID).First(&order).Error

	fmt.Println("Err2 :", err)
	if err != nil {
		res.Meta.Message = "Not Found, Try Again"
		return res, err
	}
	res.Meta.Code = 200
	res.Meta.Message = "Success"
	res.Meta.Status = "Success"
	res.Data = order
	return res, nil
}

func UpdateOrder(req apimodels.Request, orderId string) (apimodels.Response, error) {
	order := &models.Order{}
	order.OrderAt = currentTime()
	db := database.GetDb()
	err := db.Preload("Items").Where("order_id = ?", orderId).First(&order).Error
	if err != nil {
		return apimodels.Response{
			ResponseCode: "-1",
			Status:       "Failed",
		}, err
	}

	var newItems []models.Item
	var total int64

	log.Println("Request : ", req)
	log.Println("Request Items : ", req.Items)

	for _, itemreq := range req.Items {
		log.Println("ItemReq : ", itemreq)
		for _, vitem := range order.DetaiItem {
			if vitem.ID == itemreq.ID {
				log.Println("Found")
				vitem.Price = itemreq.Price
				vitem.Price = itemreq.Price
				vitem.Quantity = int(itemreq.Quantity)
				vitem.ItemCode = itemreq.ItemCode
				vitem.Description = itemreq.Description

				newItems = append(newItems, vitem)
				break
			}
		}
		total += (int64(itemreq.Quantity) * itemreq.Price)
	}

	order.DetaiItem = newItems
	errdb := db.Save(&order).Error
	if errdb != nil {
		return apimodels.Response{
			ResponseCode: "-1",
			Status:       "Failed",
		}, err
	}

	return apimodels.Response{
		Data:         req,
		DateTrans:    fmt.Sprintf("%v", dateTimeEpoch(currentTime())),
		OrderID:      order.OrderID,
		ResponseCode: "00",
		Status:       "Success",
		Total:        total,
	}, nil

}

func DeleteOrder(orderID string) bool {
	db := database.GetDb()
	order := models.Order{
		OrderID: orderID,
	}
	err := db.Where(&order).Delete(&order).Error
	if err != nil {
		log.Println("Error : ", err)
		log.Println("DeleteOrder failed")
		return false
	}

	return true

}

func generateRandomID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return fmt.Sprintf("%v", currentTime())
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// epoch time
func currentTime() int64 {
	return time.Now().Unix()
}

func dateTimeEpoch(epoch int64) time.Time {
	return time.Unix(epoch, 0)
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
