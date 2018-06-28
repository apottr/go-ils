package main

import (
	"fmt"
	"log"
	"net/http"
)

func renderJS(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s /bundle", r.Method)
	http.ServeFile(w, r, "js/dist/bundle")
}

func renderPage(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s /", r.Method)
	http.ServeFile(w, r, "html/index.html")
}

func handleMARC(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s /marc", r.Method)
	query := r.URL.Query()
	if r.Method == "GET" {
		if _, ok := query["upc"]; ok {
			r := createEmptyMARC(r.URL.Query().Get("upc"))
			err := saveMARC(r)
			if err != nil {
				log.Println(err)
			}
		} else if len(query) == 0 {
			s, err := getMARCRecords()
			if err != nil {
				log.Println(err)
			}
			fmt.Fprint(w, s)
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
