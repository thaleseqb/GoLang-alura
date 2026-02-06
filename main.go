package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Pizza struct {
	id    uuid.UUID
	name  string
	price float64
}

func main() {
	var pizzas = []Pizza{
		{id: uuid.New(), name: "toscana", price: 49.9},
		{id: uuid.New(), name: "marguerita", price: 30},
		{id: uuid.New(), name: "quatro queijos", price: 45},
	}
	for _, value := range pizzas {
		fmt.Println("Pizza cujo nome é", value.name)
	}
}
