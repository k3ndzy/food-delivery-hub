package delivery

import "testing"

func TestDeliveryService_PlaceOrder(t *testing.T) {
	svc := NewDeliveryService()
	ord, err := svc.PlaceOrder("Tuấn Anh", "123 Cầu Giấy, Hà Nội", []int{1, 2})
	if err != nil {
		t.Fatalf("lỗi đặt hàng: %v", err)
	}

	expectedSubtotal := 95000.0
	if ord.Subtotal != expectedSubtotal {
		t.Errorf("kỳ vọng subtotal %.2f nhưng nhận %.2f", expectedSubtotal, ord.Subtotal)
	}

	if ord.ShippingFee != 20000.0 {
		t.Errorf("kỳ vọng phí ship 20000 nhưng nhận %.2f", ord.ShippingFee)
	}
}
