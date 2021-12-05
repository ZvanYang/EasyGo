package middleware

import (
	"fmt"
	"net/http"
)

func AboutEndpointHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page, ")
	fmt.Fprintf(w, "where you can find information about us.")
}
