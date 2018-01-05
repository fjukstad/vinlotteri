package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var indexTemplate = template.Must(template.ParseFiles("views/index.html"))
	err := indexTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/public/", http.FileServer(http.Dir(".")))

	mux.HandleFunc("/", IndexHandler)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	fmt.Println("Server started on port", port)
	err := http.ListenAndServe(":"+port, mux)

	if err != nil {
		fmt.Println(err)
		return
	}

}
