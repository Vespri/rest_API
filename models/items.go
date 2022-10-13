package models

type Items struct {
	Item_id     string `json:"item_id"`
	Item_code   string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Order_id    string `json:"order_id"`
}
