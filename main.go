package main

import (
	"net/http"
	"strconv"

	"github.com/carvalhocaio/go-api-rest/models"
	"github.com/gin-gonic/gin"
)

var pizzas = []models.Pizza{
	{ID: 1, Nome: "toscana", Preco: 23.5},
	{ID: 2, Nome: "marguerita", Preco: 27.3},
}

func getPizzas(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"pizzas": pizzas})
}

func postPizzas(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	pizzas = append(pizzas, newPizza)
}

func getPizzasByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	for _, p := range pizzas {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Pizza not found"})
}

func main() {
	r := gin.Default()
	r.GET("/pizzas", getPizzas)
	r.POST("/pizzas", postPizzas)
	r.GET("/pizzas/:id", getPizzasByID)

	r.Run()
}
