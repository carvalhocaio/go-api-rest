package main

import (
	"net/http"

	"github.com/carvalhocaio/go-api-rest/models"
	"github.com/gin-gonic/gin"
)

func getPizzas(c *gin.Context) {
	var pizzas = []models.Pizza{
		{ID: 1, Nome: "toscana", Preco: 23.5},
		{ID: 2, Nome: "marguerita", Preco: 27.3},
	}

	c.JSON(http.StatusOK, gin.H{"pizzas": pizzas})
}

func main() {
	r := gin.Default()
	r.GET("/pizzas", getPizzas)

	r.Run()
}
