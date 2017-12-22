package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var content []string

func init() {
	content = append(content, `Welcome to the chalkboard

edit this text

right now the data is ephemeral, just like a real chalkboard ;p
`)
}

func main() {
	fmt.Println("Hello chalkboard!")

	http.HandleFunc("/api/board", apiHandler)

	static := http.FileServer(http.Dir("static"))
	http.Handle("/", static)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		log.Println(string(body))
		content = append(content, string(body))
		w.WriteHeader(http.StatusOK)
		return

	case http.MethodGet:
		fmt.Fprintf(w, content[len(content)-1])

	}
}
