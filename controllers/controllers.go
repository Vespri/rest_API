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
