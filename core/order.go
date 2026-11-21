package delivery

import "time"

type Dish struct {
	ID          int     `json:"id"`
	RestaurantID int    `json:"restaurant_id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Rating      float64 `json:"rating"`
}

type DeliveryOrder struct {
	ID             int       `json:"id"`
	CustomerName   string    `json:"customer_name"`
	Address        string    `json:"address"`
	DishIDs        []int     `json:"dish_ids"`
	Subtotal       float64   `json:"subtotal"`
	ShippingFee    float64   `json:"shipping_fee"`
	GrandTotal     float64   `json:"grand_total"`
	Status         string    `json:"status"` // PREPARING, DELIVERING, DELIVERED
	CreatedAt      time.Time `json:"created_at"`
}
