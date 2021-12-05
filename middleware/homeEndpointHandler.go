package middleware

import (
	"fmt"
	"net/http"
)

func HomeEndpointHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love Go! This is the home page.")
}
