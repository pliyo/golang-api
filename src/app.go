package main

import (
	"fmt"
	"log"
	"net/http"
)

type Api struct {
	Address string
}

func (a *Api) Run() {
	fmt.Print("Initialising routes...\n")
	router := NewRouter()
	fmt.Print("Running on ", a.Address, "\n")

	log.Fatal(http.ListenAndServe(a.Address, router))
}
