package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/carvalhocaio/go-api-rest/models"
	"github.com/gin-gonic/gin"
)

var pizzas []models.Pizza

func getPizzas(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"pizzas": pizzas})
}

func postPizzas(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	newPizza.ID = len(pizzas) + 1
	pizzas = append(pizzas, newPizza)
	savePizza()
	c.JSON(http.StatusCreated, newPizza)
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

func loadPizzas() {
	file, err := os.Open("dados/pizzas.json")
	if err != nil {
		fmt.Println("Error file:", err)
		return
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pizzas); err != nil {
		fmt.Println("Error decoding JSON: ", err)
	}
}

func savePizza() {
	file, err := os.Create("dados/pizzas.json")
	if err != nil {
		fmt.Println("Error file:", err)
		return
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(pizzas); err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}

func deletePizzaById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	for i, p := range pizzas {
		if p.ID == id {
			pizzas = append(pizzas[:i], pizzas[i+1:]...)
			savePizza()
			c.JSON(http.StatusOK, gin.H{"message": "pizza deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "pizza not found"})
}

func updatePizzaById(c *gin.Context) {
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

	for i, p := range pizzas {
		if p.ID == id {
			pizzas[i] = updatedPizza
			pizzas[i].ID = id
			savePizza()
			c.JSON(http.StatusOK, pizzas[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "pizza not found"})
}

func main() {
	loadPizzas()
	r := gin.Default()
	r.GET("/pizzas", getPizzas)
	r.POST("/pizzas", postPizzas)
	r.GET("/pizzas/:id", getPizzasByID)
	r.DELETE("/pizzas/:id", deletePizzaById)
	r.PUT("/pizzas/:id", updatePizzaById)

	r.Run()
}
