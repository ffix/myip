package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const defaultListenAddr string = ":8000"

func IpHandler(writer http.ResponseWriter, request *http.Request) {
	if ip, _, err := net.SplitHostPort(request.RemoteAddr); err == nil {
		fmt.Fprintln(writer, ip)
		return
	}
	if net.ParseIP(request.RemoteAddr) != nil {
		fmt.Fprintln(writer, request.RemoteAddr)
		return
	}
	http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	return
}

func getArgs() (string, bool) {
	listenAddr := flag.String("p", defaultListenAddr, "Listen address:port")
	useRealIp := flag.Bool("r", false, "Use RealIP header")
	flag.Parse()
	return *listenAddr, *useRealIp
}

func main() {
	listenAddr, useRealIp := getArgs()
	log.Printf("Starting server on %s...\n", listenAddr)
	router := chi.NewRouter()
	if useRealIp {
		router.Use(middleware.RealIP)
	}
	router.Use(middleware.Logger)
	router.Get("/", IpHandler)
	log.Fatal(http.ListenAndServe(listenAddr, router))
}
