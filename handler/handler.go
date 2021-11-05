package handler

import (
	"golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

//note : kapital ex HomeHan.. bersifat global

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path) //print route berjalan di cmd

	if r.URL.Path != "/" { //agar mengarah 404 jika route nya beda
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Home"))
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		http.Error(w, "error is happening", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "error is happening", http.StatusInternalServerError)
		return
	}
	// data := map[string]interface{}{
	// 	"title":   "im learn golang",
	// 	"content": "im learn golang to",
	// }

	//data := entity.Product{ID: 1, Name: "eskrim", Price: 2000, Stock: 3}

	data := []entity.Product{
		{ID: 1, Name: "es fucekbuy modal boba", Price: 20000, Stock: 3},
		{ID: 2, Name: "es krim mokodo", Price: 2000, Stock: 100},
		{ID: 3, Name: "es krim cowo baek2", Price: 100000, Stock: 1},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "error is hapening, keep calm", http.StatusInternalServerError)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World, saya sedang belajar golang web/ cara 1"))
}
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return

	}
	// w.Write([]byte("product page"))

	// fmt.Fprintf(w, "product page: %d", idNumb)
	data := map[string]interface{}{
		"content": idNumb,
	}
	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		http.Error(w, "error is happening", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "error is hapening, keep calm", http.StatusInternalServerError)
		return
	}
}

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("ini adalah GET"))
	case "POST":
		w.Write([]byte("ini adalah POST"))
	default:
		http.Error(w, "Error is happening, keep calm", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "error is happening, keep calm", http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, "error is happening, keep calm", http.StatusBadRequest)
}

func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		message := r.Form.Get("name")

		// w.Write([]byte(name))
		data := map[string]interface{}{
			"name":    name,
			"message": message,
		}
		tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "result.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "error is happening, keep calm", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "error is happening, keep calm", http.StatusInternalServerError)
			return
		}
		return
	}

	http.Error(w, "error is hapening, keep calm", http.StatusBadRequest)
}
