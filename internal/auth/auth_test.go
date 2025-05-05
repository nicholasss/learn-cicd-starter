package auth

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestNormalGetAPIKey(t *testing.T) {
	type test struct {
		input     http.Header
		expecting string
	}

	tests := []test{
		{
			input: http.Header{
				"Authorization": []string{"ApiKey token123"},
			},
			expecting: "token123",
		},
		{
			input: http.Header{
				"Authorization": []string{"ApiKey k3had987ghd23"},
			},
			expecting: "k3had987ghd23",
		},
		{
			input: http.Header{
				"Authorization": []string{"ApiKey hnva087ag63as"},
			},
			expecting: "hnva087ag63as",
		},
	}

	for _, test := range tests {
		output, err := GetAPIKey(test.input)
		if err != nil {
			t.Errorf("GetAPIKey function returned error: %q", err)
		}

		if test.expecting != output {
			t.Errorf("Expected: %q, Got: %q", test.expecting, output)
		}
	}
}

func TestFailedGetAPIKey(t *testing.T) {
	type test struct {
		input     http.Header
		expecting error
	}

	tests := []test{
		{
			input: http.Header{
				"Authorization": []string{"token123"},
			},
			expecting: errors.New("malformed authorization header"),
		},
		{
			input: http.Header{
				"Authorization": []string{"apikey uroaishdf"},
			},
			expecting: errors.New("malformed authorization header"),
		},
		{
			input: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expecting: errors.New("malformed authorization header"),
		},
	}

	for _, test := range tests {
		_, err := GetAPIKey(test.input)
		if errors.Is(test.expecting, err) && err != nil {
			fmt.Println("error from test:", err)
			t.Errorf("Expected: %q, Got: %q", test.expecting, err.Error())
		}
	}
}

// func TestEmptyGetAPIKey(t *testing.T) {
// 	type test struct {
// 		input     http.Header
// 		expecting error
// 	}
//
// 	tests := []test{
// 		{
// 			input: http.Header{
// 				"Authorization": []string{""},
// 			},
// 			expecting: errors.New("no authorization header included"),
// 		},
// 	}
//
// 	for _, test := range tests {
// 		_, err := GetAPIKey(test.input)
// 		if errors.Is(test.expecting, err) && err != nil {
// 			fmt.Println("error from test:", err)
// 			t.Errorf("Expected: %q, Got: %q", test.expecting, err.Error())
// 		}
// 	}
// }
