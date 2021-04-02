package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/thinkofher/horror"
)

// horror is pretty simple package. You can use it for building more
// complicated systems. It's very popular for REST APIs to have one
// standard implementation of error response. Below I'll try to write one.

type jsonError struct {
	Code   int    `json:"code"`
	Reason string `json:"reason"`
}

func (e jsonError) Error() string {
	return fmt.Sprintf("json error: code=%d, reason=%s", e.Code, e.Reason)
}

func (e jsonError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(e.Code)
	_ = json.NewEncoder(w).Encode(e)
}

// square is horror.HandlerFunc that accepts number through arg url query
// parameter and returns result of square of given number in form of json
// response.
func square(w http.ResponseWriter, r *http.Request) error {
	input := r.URL.Query().Get("arg")

	if input == "" {
		return jsonError{
			Code:   http.StatusBadRequest,
			Reason: "invalid input, please enter some number through arg url param",
		}
	}

	parsedInput, err := strconv.Atoi(input)
	if err != nil {
		return jsonError{
			Code:   http.StatusBadRequest,
			Reason: "invalid input, please use regular integers",
		}
	}

	var response struct {
		Answer string `json:"answer"`
	}
	response.Answer = strconv.Itoa(parsedInput * parsedInput)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		return jsonError{
			Code:   http.StatusInternalServerError,
			Reason: "internal server error, please try again later",
		}
	}

	return nil
}

func main() {
	// initialize new http.ServeMux
	mux := http.NewServeMux()

	// mount adapted handler with horror.WithError function
	mux.Handle("/square", horror.WithError(horror.HandlerFunc(square)))

	// serve and listen with mux
	log.Fatal(http.ListenAndServe(":8080", mux))

}
