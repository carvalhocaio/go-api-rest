package main

import (
	"github.com/carvalhocaio/go-api-rest/internal/data"
	"github.com/carvalhocaio/go-api-rest/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	data.LoadPizzas()

	r := gin.Default()
	r.GET("/pizzas", handler.GetPizzas)
	r.POST("/pizzas", handler.PostPizzas)
	r.GET("/pizzas/:id", handler.GetPizzasByID)
	r.DELETE("/pizzas/:id", handler.DeletePizzaById)
	r.PUT("/pizzas/:id", handler.UpdatePizzaById)

	r.Run()
}
