package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

func main() {

	cert, _ := tls.LoadX509KeyPair("./certs/domain.crt", "./certs/domain.key")

	s := &http.Server{
		Addr:    ":9000",
		Handler: nil, // use `http.DefaultServeMux`
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World!")
	})

	log.Fatal(s.ListenAndServeTLS("", ""))

}
