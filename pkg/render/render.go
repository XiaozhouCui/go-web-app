package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders templates, tmpl is the file name (e.g. home.page.tmpl)
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// get template cache from app config
	tc, err := CreateTemplateCache()
	if err != nil {
		// if there is an error, kill the app
		log.Fatal(err) // this includes os.Exit(1)
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		// if template not found, kill the app
		log.Fatal(err)
	}

	buf := new(bytes.Buffer) // create buffer to hold bytes

	_ = t.Execute(buf, nil) // write template to buffer

	// render the template
	_, err = buf.WriteTo(w) // write from buffer to http response
	if err != nil {
		log.Println("Error writing template to browser", err)
	}
}

// Cache templates to avoid loading template files from disk on every request
// it returns a map (myCache) and an error
func CreateTemplateCache() (map[string]*template.Template, error) {
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

		// find all layout-templates
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
