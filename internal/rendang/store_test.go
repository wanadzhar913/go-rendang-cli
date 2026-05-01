package rendang

import "testing"

func TestPlaceOrderReducesRemainingStock(t *testing.T) {
	store := NewStore("Test Event", 10)

	summary, err := store.PlaceOrder(Order{
		FirstName: "Faiq",
		LastName:  "Adzlan",
		Email:     "faiq@example.com",
		Packs:     4,
	})
	if err != nil {
		t.Fatalf("PlaceOrder() returned unexpected error: %v", err)
	}

	if summary.RemainingStock != 6 {
		t.Fatalf("expected remaining stock 6, got %d", summary.RemainingStock)
	}

	if summary.OrderCount != 1 {
		t.Fatalf("expected order count 1, got %d", summary.OrderCount)
	}
}

func TestPlaceOrderRejectsInvalidOrders(t *testing.T) {
	store := NewStore("Test Event", 5)

	_, err := store.PlaceOrder(Order{
		FirstName: "F",
		LastName:  "A",
		Email:     "not-an-email",
		Packs:     0,
	})
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}
}

func TestPlaceOrderRejectsOrdersLargerThanRemainingStock(t *testing.T) {
	store := NewStore("Test Event", 3)

	_, err := store.PlaceOrder(Order{
		FirstName: "Faiq",
		LastName:  "Adzlan",
		Email:     "faiq@example.com",
		Packs:     4,
	})
	if err == nil {
		t.Fatal("expected stock error, got nil")
	}
}
