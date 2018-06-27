package main

import (
	"log"
	"net/http"
)

func renderJS(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "js/dist/bundle")
}

func renderPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/index.html")
}

func handleMARC(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s /marc", r.Method)
	if r.Method == "GET" {
		r := createEmptyMARC(r.URL.Query().Get("upc"))
		err := saveMARC(r)
		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	port := ":8080"
	http.HandleFunc("/bundle", renderJS)
	http.HandleFunc("/", renderPage)
	http.HandleFunc("/marc", handleMARC)
	log.Printf("Listening on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
