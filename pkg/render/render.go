package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate renders templates using html/templates
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

// create a map to cache templates, as loading template files from disk is expensive
var tc = make(map[string]*template.Template)

// RenderTemplate renders templates using html/templates
func RenderTemplate(w http.ResponseWriter, t string) {
	// t is the template file name, e.g. home.page.tmpl
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in our cache
	_, inMap := tc[t]
	if !inMap {
		// need to create the template
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have the tempalte in the cache
		log.Println("using cached template")
	}

	tmpl = tc[t]

	// execute the cached template
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	// t is the template file name, e.g. home.page.tmpl
	// create a slice of string
	templates := []string{
		fmt.Sprintf("./templates/%s", t), // Sprintf returns the string
		"./templates/base.layout.tmpl",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to cache (map)
	tc[t] = tmpl

	return nil
}
