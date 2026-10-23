package main

import (
	"errors"
	"fmt"
)

type Item struct {
	Name  string
	Price int
}

type Order struct {
	ID     int
	Items  []Item
	Status string
}

var statuses = map[string]string{
	"new":       "Новый",
	"paid":      "Оплачен",
	"delivered": "Доставлен",
}

func NewOrder(id int) (Order, error) {
	if id <= 0 {
		return Order{}, errors.New("invalid id")
	}
	return Order{
		ID:     id,
		Status: "new",
	}, nil
}

func AddItem(order *Order, name string, price int) error {
	if order == nil {
		return errors.New("order is nil")
	}
	if name == "" {
		return errors.New("empty name")
	}
	if price < 0 {
		return errors.New("negative price")
	}
	order.Items = append(order.Items, Item{
		Name:  name,
		Price: price,
	})
	return nil
}

func Total(order Order) int {
	sum := 0
	for _, item := range order.Items {
		sum += item.Price
	}
	return sum
}

func main() {
	order, err := NewOrder(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	var name string
	var price int
	for {
		fmt.Println("Введите название товара (0 - закончить):")
		fmt.Scan(&name)
		if name == "0" {
			break
		}
		fmt.Println("Введите цену:")
		fmt.Scan(&price)
		err = AddItem(&order, name, price)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Сумма заказа:", Total(order))
}
