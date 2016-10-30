package view

import (
	"./plugin"
	"html/template"
	"net/http"
)

var (
	extensionFile = ".html"
	rootPath      = "/Users/blythe/Dropbox/Repository/Golang/Technicle/public/"
)

type View struct {
	Name    string
	Folder  string
	Vars    interface{}
	request *http.Request
}

func New(req *http.Request) *View {
	v := &View{}
	v.request = req

	return v
}

func (v *View) Render(w http.ResponseWriter) {
	t := template.Must(template.New(v.Folder).Funcs(plugin.ToUpper()).ParseFiles(rootPath+"tmpl/header.html", rootPath+"tmpl/footer.html", rootPath+v.Folder+extensionFile))

	if err := t.ExecuteTemplate(w, v.Name+extensionFile, v.Vars); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
