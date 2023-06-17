package main

import "fmt"

type item struct {
	id       string
	quantity int
	price    float64
}

type order struct {
	id    int
	items []item
}

func (order order) totalOrder() float64 {
	total := 0.0
	for _, item := range order.items {
		total += item.price * float64(item.quantity)
	}
	return total
}

func main() {

	order1 := order{id: 10001, items: []item{
		{id: "uuid1234", quantity: 10, price: 3.00},
		{id: "uuid1222", quantity: 5, price: 5.00},
		{id: "uuid1299", quantity: 2, price: 10.00},
		{id: "uuid1277", quantity: 1, price: 500.00},
		{id: "uuid1235", quantity: 2, price: 9.00},
	}}

	fmt.Println(order1.totalOrder())

}
