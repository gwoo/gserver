package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var port = flag.Int("port", 8888, "Server port.")

func main() {
	flag.Parse()
	path := flag.Arg(0)
	http.Handle("/", http.FileServer(http.Dir(path)))
	log.Printf("Serving %s on port %d\n", path, *port)

	_, cerr := os.Open("cert.pem")
	_, kerr := os.Open("key.pem")

	if os.IsNotExist(cerr) || os.IsNotExist(kerr) {
		http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
		return
	}
	log.Printf("SSL enabled.")
	http.ListenAndServeTLS(fmt.Sprintf(":%d", *port), "cert.pem", "key.pem", nil)
}
