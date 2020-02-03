package chapter12

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

// This appears as examples in godoc. Run godoc -http :8000 and hit godoc throw browser.

// ExampleSendJson provides a basic example of server.
func ExampleSendJson() {

	http.HandleFunc("/sendJson", SendJson)
	r := httptest.NewRequest("GET", "/sendJson", nil)
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, r)

	var u struct {
		Name  string
		Email string
	}

	if err := json.NewDecoder(w.Body).Decode(&u); err != nil {
		log.Println("ERROR", err)
	}

	fmt.Println(u)
	// Output:
	// {vrs abc@abc.com}
}
