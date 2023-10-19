package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/dabkoa/golang-server/data"
)

var templates = template.Must(template.ParseFiles("resources/templates/edit.html", "resources/templates/view.html"))

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := data.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusNotFound)
	}
	log.Printf("Displaying \"%s\"", title)
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := data.LoadPage(title)
	if err != nil {
		p = &data.Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &data.Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Saving %s", title)
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, templ string, p *data.Page) {
	err := templates.ExecuteTemplate(w, templ+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
