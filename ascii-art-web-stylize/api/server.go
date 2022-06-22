package api

import (
	"ascii-art-web/app"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.html"))
}

func Server() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/ascii-art", processor)
	mux.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))
	fmt.Print("Started server at http://localhost:3000/ ")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		fmt.Println(err)
		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusNotFound) // 404
		return
	}
	if err := tpl.ExecuteTemplate(w, "index.html", nil); err != nil {
		fmt.Println(err)
		return
	}
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" || r.Method == "GET" {
		http.Redirect(w, r, "/", http.StatusNotFound) // 404
		return
	}
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusMethodNotAllowed) // 405
		return
	}
	r.ParseForm()
	input, ok := r.Form["input"]
	if !ok {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400
		return
	}
	asciiText := strings.Join(input, "")
	if !app.CheckString(asciiText) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400
		return
	}
	fonts, ok := r.Form["radioButton"]
	if !ok {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400
		return
	}
	asciiStyle := strings.Join(fonts, "")

	if (asciiStyle != "standard") && (asciiStyle != "shadow") && (asciiStyle != "thinkertoy") {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400
		return
	}

	asciiArt, err := app.First(asciiText, asciiStyle)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError) // 500
		return
	}

	data := struct {
		Ascii string
		Font  string
		Input string
	}{
		Ascii: asciiArt,
		Input: asciiText,
		Font:  asciiStyle,
	}

	fmt.Println("----------------")
	fmt.Printf("Input: %#v\nFont: %#v\nAscii:\n%s\n", data.Input, data.Font, data.Ascii)
	fmt.Println(asciiStyle)

	tpl.ExecuteTemplate(w, "index.html", data)
}
