package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
HTTP server: easy to start one

To stop: had to kill process; must be a less janky way
*/

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/Hello", handler2)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Welcome to the Internet!")
}

func handler2(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Ayyyyyloha!")
}
