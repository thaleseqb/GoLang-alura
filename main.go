package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Pizza struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price float64   `json:"price"`
}

func find_all(c *gin.Context) {
	var pizzas = []Pizza{
		{ID: uuid.New(), Name: "toscana", Price: 49.9},
		{ID: uuid.New(), Name: "marguerita", Price: 30},
		{ID: uuid.New(), Name: "quatro queijos", Price: 45},
	}
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}

func pizzas_router(router *gin.Engine, path string) {
	router.GET(path, find_all)
}

func main() {
	router := gin.Default()
	pizzas_router(router, "/pizzas")
	router.Run()
}
