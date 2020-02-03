package chapter12

import (
	"net/http"
	"testing"
)

func TestDownload(t *testing.T) {
	url := "https://www.google.com/search?q=sachin"
	statusCode := 200
	t.Log("Given testing to see the internet connectivity.")
	{
		t.Logf("\t Test 0: \t When checking %q for status code %d", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatalf("\t\t\t Failed to make the call %v", err)
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
