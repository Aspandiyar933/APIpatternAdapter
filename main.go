package main

import (
	"github.com/Aspandiyar933/APIpatternAdapter/api"
	"github.com/Aspandiyar933/APIpatternAdapter/store"
)

func main() {
	api := api.NewAPIServer(":3000", store)
	api.Serve()
}