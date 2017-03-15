package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	title := r.URL.Path[len("/view"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)

}

func main() {
	fmt.Println("Start webserver ...")
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":8080", nil)
}

