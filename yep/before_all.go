package yep

import (
	"net/http"
	"text/template"
)

func Interesting(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("waterloo.html")
	t.Execute(w, nil)
}
