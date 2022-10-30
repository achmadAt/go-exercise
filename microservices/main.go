package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Println("hello world")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		res := string(data[:])
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("hello %s\n", res)
		// fmt.Printf("w: %v\n", w)
	})
	log.Println("server")
	err := http.ListenAndServe(":8000", nil)
	log.Fatal(err)
}
