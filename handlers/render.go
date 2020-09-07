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

// RenderMainPage the News Page
func (n *News) RenderMainPage(w http.ResponseWriter, r *http.Request) {
	n.l.Info("/GET mainPage")
	Arr, err := n.FetchNewsHeadlines(w, r)
	if err != nil {
		n.l.Error(err)
	}
	n.l.Info(Arr[0].Title)
	err = tpl.ExecuteTemplate(w, "news.html", Arr[0:5])
	if err != nil {
		n.l.Error(err)
	}
}

