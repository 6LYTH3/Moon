package controller

import (
	"../shared/view"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Model struct {
	Title string
	Name  string
}

func Home(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	v := view.New(r)
	v.Name = "home"
	v.Folder = "blog/home"

	v.Vars = Model{Title: "Home"}
	v.Render(w)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	v := view.New(r)
	v.Name = "hello"
	v.Folder = "blog/hello"

	// v.Vars = Model{Name: ps.ByName("name")}
	v.Vars = Model{Name: r.FormValue("name")}
	v.Render(w)
}
