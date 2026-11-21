package delivery

import (
	"errors"
	"sync"
	"time"
)

type DeliveryService struct {
	mu     sync.RWMutex
	dishes map[int]*Dish
	orders map[int]*DeliveryOrder
	lastID int
}

func NewDeliveryService() *DeliveryService {
	s := &DeliveryService{
		dishes: make(map[int]*Dish),
		orders: make(map[int]*DeliveryOrder),
	}
	s.dishes[1] = &Dish{ID: 1, RestaurantID: 101, Name: "Bún Bò Huế Chả Cua", Price: 60000, Rating: 4.8}
	s.dishes[2] = &Dish{ID: 2, RestaurantID: 101, Name: "Bánh Mì Kẹp Thịt Nướng", Price: 35000, Rating: 4.9}
	s.dishes[3] = &Dish{ID: 3, RestaurantID: 102, Name: "Trà Sữa Trân Châu Đường Đen", Price: 40000, Rating: 4.7}
	return s
}

func (s *DeliveryService) PlaceOrder(customer, address string, dishIDs []int) (*DeliveryOrder, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(dishIDs) == 0 {
		return nil, errors.New("giỏ hàng không được để trống")
	}

	var subtotal float64
	for _, id := range dishIDs {
		d, ok := s.dishes[id]
		if !ok {
			return nil, errors.New("món ăn không tồn tại")
		}
		subtotal += d.Price
	}

	shippingFee := 20000.0 // Phí ship cố định
	if subtotal > 150000 {
		shippingFee = 0.0 // Freeship cho đơn lớn
	}

	s.lastID++
	order := &DeliveryOrder{
		ID:           s.lastID,
		CustomerName: customer,
		Address:      address,
		DishIDs:      dishIDs,
		Subtotal:     subtotal,
		ShippingFee:  shippingFee,
		GrandTotal:   subtotal + shippingFee,
		Status:       "PREPARING",
		CreatedAt:    time.Now(),
	}
	s.orders[s.lastID] = order
	return order, nil
}
