package main

import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/avvooturi/car-api/pkg/handlers"
)

func main() {
	config.Load()
	store.InitStore()
	http.ListenAndServe(":3000", router)
}
