package main

import (
	"log"
	_ "net/http"

	"github.com/KornilovLN/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
