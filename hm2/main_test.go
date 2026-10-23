package main

import "testing"

func TestNewOrder(t *testing.T) {
	order, err := NewOrder(1)
	if err != nil {
		t.Fatal(err)
	}
	if Total(order) != 0 {
		t.Error("total should be 0")
	}
	_, err = NewOrder(0)
	if err == nil {
		t.Error("expected error")
	}
}

func TestAddItem(t *testing.T) {
	order, _ := NewOrder(1)
	err := AddItem(&order, "Apple", 100)
	if err != nil {
		t.Fatal(err)
	}
	err = AddItem(&order, "Milk", 200)
	if err != nil {
		t.Fatal(err)
	}
	if Total(order) != 300 {
		t.Error("total should be 300")
	}
	if len(order.Items) != 2 {
		t.Error("should have 2 items")
	}
}

func TestAddItemErrors(t *testing.T) {
	order, _ := NewOrder(1)
	if AddItem(nil, "Apple", 100) == nil {
		t.Error("expected nil error")
	}
	if AddItem(&order, "", 100) == nil {
		t.Error("expected empty name error")
	}
	if AddItem(&order, "Apple", -100) == nil {
		t.Error("expected negative price error")
	}
	if len(order.Items) != 0 {
		t.Error("items should be empty")
	}
}

func TestOrderCopy(t *testing.T) {
	order, _ := NewOrder(1)
	copyOrder := order
	copyOrder.ID = 2
	copyOrder.Status = "paid"
	if order.ID != 1 {
		t.Error("original ID changed")
	}
	if order.Status != "new" {
		t.Error("original status changed")
	}
}

func TestMap(t *testing.T) {
	var m map[string]string
	m = make(map[string]string)
	m["new"] = "Новый"
	if m["new"] != "Новый" {
		t.Error("wrong status")
	}
}
