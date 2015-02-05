package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"runtime"
	"time"
)

var (
	addr = flag.String("http", ":80", "")
	next = flag.String("next", "http://httpbin.org", "")
)

func init() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func NewServer(u *url.URL) *http.Server {
	h := httputil.NewSingleHostReverseProxy(u)
	s := &http.Server{
		Addr:           *addr,
		Handler:        h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s
}

func main() {
	u, err := url.Parse(*next)
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalln(NewServer(u).ListenAndServe())
}
