package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	gitProxyURL, err := url.Parse("http://127.0.0.1:7001")
	if err != nil {
		log.Fatal(err)
	}
	gitProxy := httputil.NewSingleHostReverseProxy(gitProxyURL)

	appProxyURL, err := url.Parse("http://127.0.0.1:5000")
	if err != nil {
		log.Fatal(err)
	}
	appProxy := httputil.NewSingleHostReverseProxy(appProxyURL)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/git/") {
			gitProxy.ServeHTTP(w, r)
			return
		}

		if r.URL.Path == "/git" {
			http.Redirect(w, r, "/git/", http.StatusPermanentRedirect)
			return
		}

		appProxy.ServeHTTP(w, r)
	})

	fmt.Println("Server listening on 127.0.0.1:7002")
	log.Fatal(http.ListenAndServe("127.0.0.1:7002", nil))
}
