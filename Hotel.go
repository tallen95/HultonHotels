package main

import(
	"fmt"
	"net/http"
	"log"
)

func roothandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	//serves the static files in directory static
	http.Handle("/", http.FileServer(http.Dir("./static")))
	//listens on port 8081
	log.Fatal(http.ListenAndServe(":8081", nil))
}
