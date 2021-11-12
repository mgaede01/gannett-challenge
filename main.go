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
	{ProduceCode: "ABCD-EFGH-IJKL-MNOP", Name: "Zucchini", UnitPrice: 1.17},
}

func getProduceList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, produceList)
}

func main() {
	fmt.Println("Starting gin...")
	router := gin.Default()
	router.GET("/", getProduceList)
	router.Run("localhost:8080")
}
