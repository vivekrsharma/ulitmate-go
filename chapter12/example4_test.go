package chapter12

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInternalHandlers(t *testing.T) {

	statusCode := 200
	url := "/sendJson"
	http.HandleFunc("/sendJson", SendJson)

	t.Log("Given testing to see the internet connectivity.")
	{

		r := httptest.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(w, r)

		t.Logf("\t Test 0: \t When checking %q for status code %d", url, statusCode)
		{

			if w.Code == statusCode {
				t.Logf("\t\t\t Should receive %d as status code", statusCode)
			} else {
				t.Errorf("\t\t\t Should receive %d error code but got %d", statusCode, w.Code)
			}

			var u struct {
				Name  string
				Email string
			}

			if err := json.NewDecoder(w.Body).Decode(&u); err != nil {
				t.Fatalf("Should be able to decode into u.")
			}

			if u.Name == "vrs" {
				t.Logf("\t\t\t Should receive vrs as Name")
			} else {
				t.Errorf("\t\t\t Should receive vrs as Name but got %s", u.Name)
			}

			if u.Email == "abc@abc.com" {
				t.Logf("\t\t\t Should receive abc@abc.com as Email")
			} else {
				t.Errorf("\t\t\t Should receive abc@abc.com as Email but got %s", u.Email)
			}
		}
	}
}
