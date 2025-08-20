package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// http://localhost:8080/double?user_id=2
func DoubleHandler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("value")

	if text == "" {
		http.Error(w, "value is empty", http.StatusBadRequest)
		return
	}

	n, err := strconv.Atoi(text)
	if err != nil {
		http.Error(w, "value is invalid", http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, n*2)

}

//func main() {
//	http.HandleFunc("/double", DoubleHandler)
//	http.ListenAndServe(":8080", nil)
//}
