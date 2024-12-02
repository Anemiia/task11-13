package main

import (
	"fmt"
)

type Dish struct {
	Name  string
	Price float64
}

type Order struct {
	Dishes []Dish
	Total  float64
}

func (order *Order) AddDish(dish Dish) {
	order.Dishes = append(order.Dishes, dish)
}

func (order *Order) RemoveDish(dish Dish) {
	for i, d := range order.Dishes {
		if d.Name == dish.Name && d.Price == dish.Price {
			order.Dishes = append(order.Dishes[:i], order.Dishes[i+1:]...) // троетоечие для распаковки слайса
			break
		}
	}
}
func (order *Order) CalculateTotal() {
	var total float64

	for _, d := range order.Dishes {
		total += d.Price
	}

	order.Total = total
}


func main() {
	order := Order{}
	dish1 := Dish{Name: "Pizza", Price: 10.99}
	dish2 := Dish{Name: "Burger", Price: 5.99}

	order.AddDish(dish1)
	order.AddDish(dish2)

	order.CalculateTotal()
	fmt.Println("Total:", order.Total)

	order.RemoveDish(dish1)

	order.CalculateTotal()
	fmt.Println("Total:", order.Total)
}