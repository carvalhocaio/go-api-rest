package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pizza struct {
	ID    int     `json:"id"`
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}

func pizzas(c *gin.Context) {
	var pizzas = []Pizza{
		{ID: 1, Nome: "toscana", Preco: 23.5},
		{ID: 2, Nome: "marguerita", Preco: 27.3},
	}

	c.JSON(http.StatusOK, gin.H{"pizzas": pizzas})
}

func main() {
	r := gin.Default()
	r.GET("/pizzas", pizzas)

	r.Run()
}
