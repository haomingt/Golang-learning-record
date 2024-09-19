package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) error {
	_, err := writer.Write([]byte(message))
	return err
}
func englishHandler(writer http.ResponseWriter, resquest *http.Request) {
	err := write(writer, "Hello Wed")
	if err != nil {
		log.Fatal(err)
	}
}
func frenchHandler(writer http.ResponseWriter, resquest *http.Request) {
	err := write(writer, "Salut Wed")
	if err != nil {
		log.Fatal(err)
	}
}
func hindiHandler(writer http.ResponseWriter, resquest *http.Request) {
	err := write(writer, "Namaste Wed")
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/namaste", hindiHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
