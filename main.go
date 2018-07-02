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
	var upc, title, typ string
	log.Printf("%s /marc", r.Method)
	query := r.URL.Query()
	if r.Method == "GET" {
		mupc, upcOK := query["upc"]
		upc = mupc[0]
		mtitle, titleOK := query["title"]
		if titleOK {
			title = mtitle[0]
		}
		mtype, typeOK := query["type"]
		if typeOK {
			typ = mtype[0]
		}
		if upcOK && !titleOK && !typeOK {
			r := createEmptyMARC(upc)
			err := saveMARC(r)
			if err != nil {
				log.Println(err)
			}
			fmt.Fprint(w, "created MARC")
		} else if titleOK && upcOK && !typeOK {
			r, err := getMARCRecord(upc)
			if err != nil {
				log.Printf("error: %s\n", err)
			}
			setNameMARC(r, title)
			fmt.Fprintf(w, "added title %s to MARC %s", title, upc)
		} else if typeOK && upcOK && !titleOK {
			r, err := getMARCRecord(upc)
			if err != nil {
				log.Println(err)
			}
			setTypeMARC(r, typ)
			fmt.Fprintf(w, "added type %s to MARC %s", typ, upc)
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
