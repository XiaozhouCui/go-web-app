package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders templates using html/templates
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache

	// get requested template from cache

	// render the template

	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

// loading template files from disk on every request is expensive, need cache
func createTemplateCache() (map[string]*template.Template, error) {
	// create an empty map to cache templates
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl") // return a slice of strings (full path)
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page) // return the filename without path
		// parse the file and save it into a template set (ts)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// find layout templates
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		// if there are layout-templates, associate them with page-templates (ts)
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl") // use "=" instead of ":="
			if err != nil {
				return myCache, err
			}
		}

		// at the end of loop, save the template into cache (e.g. myCache['home.page.tmpl'])
		myCache[name] = ts
	}

	return myCache, nil
}
