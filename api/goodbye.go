package api

import (
	"fmt"
	"net/http"
)

func GoodbyeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Goodbye from the goodbye handler!")
}