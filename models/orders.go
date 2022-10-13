package models

type Orders struct {
	Order_id      string `json:"order_id"`
	Customer_name string `json:"customer_name"`
	Ordered_at    string `json:"ordered_at"`
	Items         []Items
}

var OrderDatas = []Orders{}
