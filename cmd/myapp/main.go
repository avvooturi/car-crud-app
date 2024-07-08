package main

import (
	"fmt"

	"github.com/avvooturi/car-crud-app/pkg/api"
	"github.com/avvooturi/car-crud-app/pkg/handlers"
	"github.com/avvooturi/car-crud-app/pkg/storage"
)

func main() {
	fmt.Println("welcome to my app!")
	server := api.NewApi()
	handler := handlers.New(storage.NewArrayDB())
	server.GET("/fetch/all", handler.HandleFetchOne)

	err := server.Start(3000)
	if err != nil {
		panic(err)
	}
}
