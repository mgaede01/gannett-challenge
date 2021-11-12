package main

import (
	"fmt"
)

type ProduceItem struct {
	ProduceCode string  `json:"ProduceCode"`
	Name        string  `json:"Name"`
	UnitPrice   float64 `json:"UnitPrice"`
}

var produceList = []ProduceItem{
	{ProduceCode: "ABCD-EFGH-IJKL-MNOP", Name: "Zucchini", UnitPrice: 1.17},
}

func main() {
	fmt.Println(produceList)
}
