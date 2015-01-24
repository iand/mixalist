package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	router.Path("/hello").HandlerFunc(hello)
	router.Path("/").HandlerFunc(frontpage)

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	server.ListenAndServe()
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprint(w, "YO")
	fmt.Fprint(w, "<br>HIII")
}
func frontpage(w http.ResponseWriter, r *http.Request) {
	box, _ := rice.FindBox("templates")

	templateData, _ := box.String("frontpage.html")
	t, _ := template.New("frontpage.html").Parse(templateData)

	data := map[string]interface{}{}
	w.Header().Add("Content-Type", "text/html")
	t.Execute(w, data)

}
