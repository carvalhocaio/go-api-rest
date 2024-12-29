package handler

import (
	"net/http"
	"strconv"

	"github.com/carvalhocaio/go-api-rest/internal/data"
	"github.com/carvalhocaio/go-api-rest/internal/models"
	"github.com/carvalhocaio/go-api-rest/internal/service"
	"github.com/gin-gonic/gin"
)

func GetPizzas(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"pizzas": data.Pizzas})
}

func PostPizzas(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := service.ValidatePizzaPrice(&newPizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	newPizza.ID = len(data.Pizzas) + 1
	data.Pizzas = append(data.Pizzas, newPizza)
	data.SavePizza()
	c.JSON(http.StatusCreated, newPizza)
}

func GetPizzasByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	for _, p := range data.Pizzas {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Pizza not found"})
}

func DeletePizzaById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	for i, p := range data.Pizzas {
		if p.ID == id {
			data.Pizzas = append(data.Pizzas[:i], data.Pizzas[i+1:]...)
			data.SavePizza()
			c.JSON(http.StatusOK, gin.H{"message": "pizza deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "pizza not found"})
}

func UpdatePizzaById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	var updatedPizza models.Pizza
	if err := c.ShouldBindJSON(&updatedPizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := service.ValidatePizzaPrice(&updatedPizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	for i, p := range data.Pizzas {
		if p.ID == id {
			data.Pizzas[i] = updatedPizza
			data.Pizzas[i].ID = id
			data.SavePizza()
			c.JSON(http.StatusOK, data.Pizzas[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "pizza not found"})
}
