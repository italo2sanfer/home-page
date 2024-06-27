package main

import (
	"log"
    "html/template"
	"net/http"
	"os"
)

type ContactDetails struct {
    Name    string
	LinkGithub string
}

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		link_github := os.Getenv("LINK_GITHUB")
		if link_github == "" {
			link_github = "https://github.com/italo2sanfer/"
		}
		details := ContactDetails{
			Name:   "Aaaa",
			LinkGithub: link_github,
		}
		tmpl.Execute(w, details)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}