package main

import (
	"github.com/anmolbabu/product_service/pkg/api"
)

func main() {
	api.StartRestServer()
	defer api.Cleanup()
}
