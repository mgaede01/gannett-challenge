package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProduceItem struct {
	ProduceCode string  `json:"ProduceCode"`
	Name        string  `json:"Name"`
	UnitPrice   float64 `json:"UnitPrice"`
}

var produceList = []ProduceItem{
	{ProduceCode: "A12T-4GH7-QPL9-3N4M", Name: "Lettuce", UnitPrice: 3.46},
	{ProduceCode: "E5T6-9UI3-TH15-QR88", Name: "Peach", UnitPrice: 2.99},
	{ProduceCode: "YRT6-72AS-K736-L4AR", Name: "Green Pepper", UnitPrice: 0.79},
	{ProduceCode: "TQ4C-VV6T-75ZX-1RMR", Name: "Gala Apple", UnitPrice: 3.59},
}

func getProduceList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, produceList)
}

func postProduceItem(c *gin.Context) {
	var newItem ProduceItem

	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	produceList = append(produceList, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}

func main() {
	fmt.Println("Starting gin...")
	router := gin.Default()
	router.GET("/", getProduceList)
	router.POST("/", postProduceItem)
	router.Run("localhost:8080")
}
