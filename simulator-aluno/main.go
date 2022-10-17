package main

import (
	"fmt"

	"github.com/h-buzzi/delivery-imersaofsfc2/application/route"
)

func main() {
	route := route.Route{
		ID:       "1",
		ClientID: "1",
	}
	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println((stringjson[1]))
}
