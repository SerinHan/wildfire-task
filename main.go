package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", returnJoke)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func fetch(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}

func returnJoke(w http.ResponseWriter, _ *http.Request) {
	fetch("https://names.mcquay.me/api/v0")
	fetch("http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=nerdy")
}