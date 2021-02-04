package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %q\n", err)
	}
}

type proxyList struct {
	proxy230 *httputil.ReverseProxy
	proxy232 *httputil.ReverseProxy
}

// Build with >> go build -o bin/sci-version-proxy -v .
func main() {
	port := os.Getenv("PORT")
	remote230, urlErr := url.Parse(os.Getenv("REMOTE_230"))
	remote232, urlErr := url.Parse(os.Getenv("REMOTE_232"))
	if urlErr != nil {
		panic(urlErr)
	}

	proxies := proxyList{
		proxy230: httputil.NewSingleHostReverseProxy(remote230),
		proxy232: httputil.NewSingleHostReverseProxy(remote232),
	}

	http.HandleFunc("/", handler(&proxies))

	log.Printf("Golang App running...\n")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(pxs *proxyList) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		apiVersion := strings.Split(r.Header.Get("X-SCI-VERSION"), ".")[0]
		apiNum, err := strconv.Atoi(apiVersion)

		if err != nil {
			log.Println("no header = 232")
			pxs.proxy232.ServeHTTP(w, r)
			return
		}

		if apiNum >= 232 {
			log.Println("232")
			pxs.proxy232.ServeHTTP(w, r)
			return
		} else if apiNum >= 230 {
			log.Println("230")
			pxs.proxy230.ServeHTTP(w, r)
			return
		}
		log.Println("Default = 232")
		pxs.proxy232.ServeHTTP(w, r)
	}
}
