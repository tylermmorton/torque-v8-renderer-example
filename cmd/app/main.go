package main

import (
	"github.com/tylermmorton/torque"
	"github.com/tylermmorton/torque-v8-renderer-example/app/routes"
	"log"
	"net/http"
)

func main() {
	h, err := torque.New[routes.ViewModel](&routes.Controller{Dist: nil})
	if err != nil {
		log.Fatalf("failed to create torque app: %v", err)
	}

	log.Printf("Starting torque app at http://localhost:8080")
	err = http.ListenAndServe(":8080", h)
	if err != nil {
		log.Fatalf("failed to start http listener: %v", err)
	}
}
