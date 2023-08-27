package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/videos/", func(w http.ResponseWriter, r *http.Request) {
		videoPath := "videos/" + r.URL.Path[len("/videos/"):]
		http.ServeFile(w, r, videoPath)
	})

	ipAddress := "localhost" // Corrected "localHost" to "localhost"
	port := ":8080"          // Added a colon before the port number
	address := ipAddress + port

	http.ListenAndServe(address, nil)
}
