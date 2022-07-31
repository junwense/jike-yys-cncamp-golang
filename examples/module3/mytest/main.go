package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe(":8081", nil)

	//mux := http.NewServeMux()
	//mux.HandleFunc("/", rootHandler)
	//mux.HandleFunc("/healthz", healthz)
	//mux.HandleFunc("/debug/pprof/", pprof.Index)
	//mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	//mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	//mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	if err != nil {
		log.Fatal(err)
	}
}

func healthzHandler(writer http.ResponseWriter, request *http.Request) {

	log.Printf("start to healthz ")
	io.WriteString(writer, "ok")
	log.Printf("end to healthz ")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("start to rootHandler ")
	os.Setenv("VERSION", "v0.0.1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)

	for k, v := range r.Header {
		w.Header().Set(k, fmt.Sprintf("%s", v))
	}

	clientip := remoteIp(r)
	log.Printf("Success! Response code: %d", 200)
	log.Printf("Success! clientip: %s", clientip)
	io.WriteString(w, "ok")
	log.Printf("end to rootHandler ")
}

//https://blog.csdn.net/love666666shen/article/details/103196623
func remoteIp(req *http.Request) string {

	log.Printf("start to remoteIp ")
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	//ipv6处理
	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	log.Printf("end to remoteIp ")

	return remoteAddr
}
