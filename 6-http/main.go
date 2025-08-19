package main

import "net/http"

func main() {

	http.HandleFunc("/home", Home)
	http.HandleFunc("/json", SendJson)
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	// panic() // if panic happens in a http request than http can recover the panic
	go func() {
		// if panic happens in a goroutine spun by the dev, then the server would crash
		// make sure to do recovery logic
		panic("something went wrong")
	}()
	w.Write([]byte("Hello World!"))

}

func SendJson(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{name:"Raj"}`))
}
