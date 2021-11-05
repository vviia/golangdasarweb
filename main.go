package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// aboutHandler := func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("cara 2"))
	// }

	mux.HandleFunc("/", handler.HomeHandler)           //root utama lebih sering dipakai
	mux.HandleFunc("/hello", handler.HelloHandler)     //root lain
	mux.HandleFunc("/product", handler.ProductHandler) //root lain
	mux.HandleFunc("/post-get", handler.PostGet)       //root lain
	mux.HandleFunc("/form", handler.Form)              //root lain
	mux.HandleFunc("/process", handler.Process)        //root lain

	// mux.HandleFunc("/helloo", aboutHandler) //root dengan cara lain
	// mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("cara root 3"))
	// }) //root anonymouse

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// localhost:9000/static/style.css
	// style.css

	log.Println("starting web on port 9000") //server port

	err := http.ListenAndServe(":9000", mux)
	log.Fatal(err)
}
