package handlers

import (
	"net/http"
	"webThird/pkg/render"
)

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplates(writer, "home.page.tmpl")
}

func AboutHandler(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplates(writer, "about.page.tmpl")
}
