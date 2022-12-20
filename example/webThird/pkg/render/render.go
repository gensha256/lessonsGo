package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmplCache = make(map[string]*template.Template)

func RenderTemplates(writer http.ResponseWriter, str string) {
	var tmpl *template.Template
	var err error

	_, inMap := tmplCache[str]
	if !inMap {
		err = makeTemplateCache(str)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("Template in cache")
	}
	tmpl = tmplCache[str]
	err = tmpl.Execute(writer, nil)
	if err != nil {
		log.Println(err)
	}
}

func makeTemplateCache(str string) error {

	tmplPath := "/home/gen/projects/go/src/github.com/gensha256/test/example/webThird/templates/"
	templates := []string{
		fmt.Sprintf(tmplPath+"%s", str), tmplPath + "base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tmplCache[str] = tmpl
	log.Println(tmplCache)
	return nil
}
