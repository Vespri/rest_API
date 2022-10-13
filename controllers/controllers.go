package controllers

import (
	"api_go/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateData(k *gin.Context) {
	var newOrder models.Orders

	if err := k.ShouldBindJSON(&newOrder); err != nil {
		k.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newOrder.Order_id = fmt.Sprintf("%d", len(models.OrderDatas)+1)
	newOrder.Items[0].Item_id = fmt.Sprintf("%d", len(models.OrderDatas)+1)
	newOrder.Items[0].Order_id = newOrder.Order_id

	fmt.Println(newOrder)

	models.OrderDatas = append(models.OrderDatas, newOrder)

	k.JSON(http.StatusCreated, gin.H{
		"Data Order Item": newOrder,
	})
}

func UpdateData(k *gin.Context) {
	orderID := k.Param("orderID")
	condition := false

	var updateOrder models.Orders

	if err := k.ShouldBindJSON(&updateOrder); err != nil {
		k.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, order := range models.OrderDatas {
		if orderID == order.Order_id {
			condition = true
			models.OrderDatas[i] = updateOrder
			models.OrderDatas[i].Order_id = orderID
			models.OrderDatas[i].Items[0].Item_id = orderID
			models.OrderDatas[i].Items[0].Order_id = orderID
			break
		}
	}

	if !condition {
		k.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Order with id %v not found", orderID),
		})
		return
	}

	k.JSON(http.StatusOK, gin.H{
		"Message": fmt.Sprintf("Order with id %v has been updated successfully", orderID),
	})
}

func GetAllData(k *gin.Context) {
	condition := true

	if !condition {
		k.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found",
		})
		return
	}

	if len(models.OrderDatas) < 1 {
		k.JSON(http.StatusOK, gin.H{
			"Order Datas": "Order Data is Empty",
		})
	} else {
		k.JSON(http.StatusOK, gin.H{
			"Order Datas": models.OrderDatas,
		})
	}

}

func GetDataByID(k *gin.Context) {
	orderID := k.Param("orderID")
	condition := false

	var orderData models.Orders

	for i, order := range models.OrderDatas {
		if orderID == order.Order_id {
			condition = true
			orderData = models.OrderDatas[i]
			break
		}
	}

	if !condition {
		k.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Order with id %v not found", orderID),
		})
		return
	}

	k.JSON(http.StatusOK, gin.H{
		"Order": orderData,
	})
}

func DeleteData(k *gin.Context) {
	orderID := k.Param("orderID")
	condition := false

	var orderIndex int

	for i, order := range models.OrderDatas {
		if orderID == order.Order_id {
			condition = true
			orderIndex = i
			break
		}
	}

	if !condition {
		k.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Order with id %v not found", orderID),
		})
		return
	}

	copy(models.OrderDatas[orderIndex:], models.OrderDatas[orderIndex+1:])
	models.OrderDatas[len(models.OrderDatas)-1] = models.Orders{}
	models.OrderDatas = models.OrderDatas[:len(models.OrderDatas)-1]

	k.JSON(http.StatusOK, gin.H{
		"Message": fmt.Sprintf("Order with id %v has been deleted successfully", orderID),
	})
}
