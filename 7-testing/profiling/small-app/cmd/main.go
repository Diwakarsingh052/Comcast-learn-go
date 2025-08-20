package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

// curl localhost:8080/find?user_id=123
// main-> if request hits -> handlers -> models
// inuse_objects: Number of live objects currently allocated.
//
// alloc_space: Total cumulative bytes allocated (even if later freed). Useful to see total churn.
//
// alloc_objects: Total number of objects ever allocated.
func main() {
	// globally set config for log package
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	s := http.NewServeMux()
	// it would initialize the conn once at the start of the app
	setupRoutes(s)

	go func() { http.ListenAndServe(":3000", nil) }()

	http.ListenAndServe(":8080", s)
}

//http://localhost:3000/debug/pprof // go tool pprof http://localhost:3000/debug/pprof/goroutine
