package main 

import (
	"fmt"
	"net/http"
	"html/template"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Batman",
	}

	var tmpl = template.Must(template.ParseFiles(
		"views/index.html",
		"views/_header.html",
		"views/_footer.html",
	))

	var err = tmpl.ExecuteTemplate(w, "index",data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func handleAbout(w http.ResponseWriter, r *http.Request){
	var msg = "This is about"
	w.Write([]byte(msg))
}

func handleContact(w http.ResponseWriter, r *http.Request){
	var msg = "this is contact"
	w.Write([]byte(msg))
}

func handlePortofolio(w http.ResponseWriter, r *http.Request){
	var msg = "this is portofolio"
	w.Write([]byte(msg))
}


func main() {
	http.Handle("/static/", 
        http.StripPrefix("/static/", 
            http.FileServer(http.Dir("assets"))))

    http.HandleFunc("/", handleIndex)
    http.HandleFunc("/about", handleAbout)
    http.HandleFunc("/contact", handleContact)
	http.HandleFunc("/portofolio", handlePortofolio)

	var address = "localhost:9000"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}