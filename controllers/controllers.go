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
	newOrder.Items[0].Item_id = fmt.Sprintf("%d", len(models.ItemDatas)+1)
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
