package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("Starting http server...")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)
	c, python, java := true, false, "no!"
	fmt.Println(c, python, java)
	err := http.ListenAndServe(":8081", nil)
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", rootHandler)
	// mux.HandleFunc("/healthz", healthz)
	// mux.HandleFunc("/debug/pprof/", pprof.Index)
	// mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	// mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	// mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	// err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	//作业四
	w.WriteHeader(200)
	io.WriteString(w, "ok\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")
	user := r.URL.Query().Get("user")
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
	} else {
		io.WriteString(w, "hello [stranger]\n")
	}
	io.WriteString(w, "===================Details of the http request URL:============\n")
	io.WriteString(w, fmt.Sprintf(r.URL.String()))
	io.WriteString(w, "\n===================Details of the http request header:============\n")

	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
		//作业一
		for _, ve := range v {
			fmt.Println("Adding k: " + k + ", v: " + ve + " into resp header.\n")
			w.Header().Add(k, ve)
		}

	}

	//作业二
	ver := os.Getenv("VERSION")
	w.Header().Add("Version", ver)
	w.WriteHeader(200)

	//作业三
	ip := r.Header.Get("X-FORWARDED-FOR")

	if ip == "" {
		ip = "127.0.0.1"
	}

	glog.V(2).Info("glog: Http request from " + ip + ", the return status code is 200\n")
	fmt.Println("fmt: Http request from " + ip + ", the return status code is 200\n")

	return
}
