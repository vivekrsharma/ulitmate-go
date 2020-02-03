package chapter12

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "applications/xml")
		fmt.Fprint(w, "<xml/>")
	}
	return httptest.NewServer(http.HandlerFunc(f))
}

func TestDownloadMockServer(t *testing.T) {

	statusCode := 200
	server := mockServer()
	defer server.Close()

	t.Log("Given testing to see the internet connectivity.")
	{
		t.Logf("\t Test 0: \t When checking %q for status code %d", server.URL, statusCode)
		{
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatalf("\t\t\t Failed to make the call %+v", err)
			}
			t.Log("\t\t\t Should be to make the http call")
			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\t\t Should receive %d as status code", statusCode)
			} else {
				t.Errorf("\t\t\t Should receive %d error code but got %d", statusCode, resp.StatusCode)
			}
		}
	}
}
