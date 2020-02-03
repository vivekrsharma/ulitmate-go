package chapter12

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/sendJson", SendJson)
	log.Println("listener: Started: Listening on http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}

// SendJson handles the json endpoint for the http server.
func SendJson(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "vrs",
		Email: "abc@abc.com",
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}
