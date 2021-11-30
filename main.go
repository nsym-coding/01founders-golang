package main

import (
	"html/template"
	"net/http"
)

/*This var is a pointer towards template.Template that is a
pointer to help process the html.*/
var tpl *template.Template

/*This init function, once it's initialised, makes it so that each html file
in the templates folder is parsed i.e. they all get looked through once and
then stored in the memory ready to go when needed*/
func init() {
	tpl = template.Must(template.ParseGlob("templates/*gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/ascii-art", asciiart)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func asciiart(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	userBanner := r.FormValue("uBanner")
	userString := r.FormValue("uString")

	d := struct {
		Banner string
		String string
	}{
		Banner: userBanner,
		String: userString,
	}

	tpl.ExecuteTemplate(w, "ascii-art.gohtml", d)
}
