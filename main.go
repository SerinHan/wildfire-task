package main

import (
	"io"
	"log"
	"net/http"
	"encoding/json"
	"strings"
)

func main() {
	http.HandleFunc("/", returnJoke)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func fetchJSON(url string) []byte {
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

	return body
}

func parseRandomName(body []byte) (string, string) {
	var result map[string]interface{}

	json.Unmarshal(body, &result)

	firstName := result["first_name"].(string)
	lastName := result["last_name"].(string)

	return firstName, lastName
}

func parseRandomJoke(body []byte) string {
	var result map[string]interface{}

	json.Unmarshal(body, &result)

	joke := result["value"].(map[string]interface{})["joke"].(string)

	return joke
}

func returnJoke(w http.ResponseWriter, _ *http.Request) {
	firstName, lastName := parseRandomName(fetchJSON("https://names.mcquay.me/api/v0"))
	joke := parseRandomJoke(fetchJSON("http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=nerdy"))

	joke = strings.Replace(joke, "John", firstName, -1)
	joke = strings.Replace(joke, "Doe", lastName, -1)
	joke = joke + "\n"

	io.WriteString(w, joke)
}