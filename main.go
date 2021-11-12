package main

import (
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

var produceCodeRegex string = "^[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{4}$"
var produceNameRegex string = "^[a-zA-Z0-9 ]*$"

type ProduceItem struct {
	ProduceCode string  `json:"Produce Code"`
	Name        string  `json:"Name"`
	UnitPrice   float64 `json:"Unit Price"`
}

var produceList = []ProduceItem{
	{ProduceCode: "A12T-4GH7-QPL9-3N4M", Name: "Lettuce", UnitPrice: 3.46},
	{ProduceCode: "E5T6-9UI3-TH15-QR88", Name: "Peach", UnitPrice: 2.99},
	{ProduceCode: "YRT6-72AS-K736-L4AR", Name: "Green Pepper", UnitPrice: 0.79},
	{ProduceCode: "TQ4C-VV6T-75ZX-1RMR", Name: "Gala Apple", UnitPrice: 3.59},
}

func isValidProduceCode(produceCode string) bool {
	result, _ := regexp.MatchString(produceCodeRegex, produceCode)
	return result
}

func isValidProduceName(produceName string) bool {
	result, _ := regexp.MatchString(produceNameRegex, produceName)
	return result
}

func getProduceList(c *gin.Context) {
	if len(produceList) == 0 {
		c.String(http.StatusNoContent, "No Produce to display")
		return
	}
	c.IndentedJSON(http.StatusOK, produceList)
}

func postProduceItem(c *gin.Context) {
	var newItem ProduceItem

	if err := c.BindJSON(&newItem); err != nil {
		return
	}
	newItem.ProduceCode = strings.ToUpper(newItem.ProduceCode)
	if !isValidProduceCode(newItem.ProduceCode) {
		c.String(http.StatusBadRequest, "Bad Produce Code\n")
		return
	}
	if !isValidProduceName(newItem.Name) {
		c.String(http.StatusBadRequest, "Bad Produce Name\n")
		return
	}
	newItem.UnitPrice = math.Round(newItem.UnitPrice*100) / 100
	// Check if item already exists
	for _, item := range produceList {
		if newItem == item {
			c.String(http.StatusBadRequest, "Item already exists\n")
			return
		}
	}

	// Add new item
	produceList = append(produceList, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}

func removeFromList(i int) {
	produceList[i] = produceList[len(produceList)-1]
	produceList = produceList[:len(produceList)-1]
}

func deleteProduceItem(c *gin.Context) {
	var produceCode string
	index := -1

	produceCode = c.Param("produceCode")
	produceCode = strings.ToUpper(produceCode)
	if !isValidProduceCode(produceCode) {
		c.String(http.StatusBadRequest, "Bad Produce Code\n")
		return
	}

	// Search for Produce Code in list
	for i, item := range produceList {
		if item.ProduceCode == produceCode {
			index = i
		}
	}
	if index == -1 {
		c.String(http.StatusNoContent, "Item with Produce Code %s not found.\n", produceCode)
		return
	}
	removeFromList(index)
	c.String(http.StatusOK, "Item with Produce Code %s deleted.\n", produceCode)
}

func main() {
	fmt.Println("Starting gin...")
	router := gin.Default()
	router.GET("/", getProduceList)
	router.POST("/", postProduceItem)
	router.DELETE("/:produceCode", deleteProduceItem)
	router.Run("localhost:8080")
}
