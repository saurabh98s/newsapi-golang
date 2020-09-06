package handlers

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

// Init exposes the templates to serve the web page
func Init() {
	tpl = template.Must(template.ParseGlob("template/*.html"))
}

// RenderPage the News Page
func (n *News) RenderPage(w http.ResponseWriter, r *http.Request) {

	Arr, err := n.FetchNewsHeadlines(w, r)
	if err != nil {
		n.l.Error(err)
	}
	n.l.Info(Arr[0].Title)
	err = tpl.ExecuteTemplate(w, "news.html", Arr)
	if err != nil {
		n.l.Error(err)
	}
}
