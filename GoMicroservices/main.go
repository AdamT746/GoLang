package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main () {
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request){
		log.Println("Hello There")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "UhOh", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "hello %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request){
		log.Println(("Goodbye World"))
	})

	http.ListenAndServe(":5040", nil)
}