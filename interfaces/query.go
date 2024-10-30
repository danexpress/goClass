package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)

}

type database map[string]dollars

func main() {
	db := database{
		"shoe":  50,
		"socks": 5,
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
