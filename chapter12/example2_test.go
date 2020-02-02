package chapter12

import (
	"net/http"
	"testing"
)

func TestDownloadTable(t *testing.T) {

	tests := []struct {
		url string
		statusCode int
	}{{
		"https://www.google.com/search?q=sachin",
		200,
	},
		{
			"http://rss.cnn.com/rss/cnn_top_storie.rss",
			404,
		},
	}

	t.Log("Given testing to see the internet connectivity.")
	{
		for _, tt := range tests {

			t.Logf("\t Test 0: \t When checking %q for status code %d", tt.url, tt.statusCode)
			{
				resp, err := http.Get(tt.url)
				if err != nil {
					t.Fatalf("\t\t\t Failed to make the call %+v", err)
				}
				t.Log("\t\t\t Should be to make the http call")
				defer resp.Body.Close()

				if resp.StatusCode == tt.statusCode {
					t.Logf("\t\t\t Should receive %d as status code",tt.statusCode)
				} else {
					t.Errorf("\t\t\t Should receive %d error code but got %d", tt.statusCode, resp.StatusCode)
				}
			}
		}
	}
}
