package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/omeg21/project-repo/pkg/config"
	"github.com/omeg21/project-repo/pkg/models"
)

var app *config.AppConfig

//sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// RenderTemplate renders tempalate using HTML template
func RenderTemplate(w http.ResponseWriter, html string, td *models.TemplateData) {
	// get the template cache from the app config
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	//get requested template from cache
	t, ok := tc[html]

	if !ok {
		log.Fatal("could not get template from cache")
	}

	buf := new(bytes.Buffer)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	//render the template
	// var pathToTemplates = "./templates/"
	// var pathToLayout  = "./templates/"
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
	// parsedTemplate, err := template.ParseFiles(pathToTemplates+html, "./templates/base.layout.html")
	// if err != nil {
	// 	fmt.Println("error parsing template", err)
	// 	returngo
	// }
	// parsedTemplate.Execute(w, nil)

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.html from ./templates/
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}
	//range trhough all files ending with *.page.html
	for _, page := range pages {
		//get file name
		fileName := filepath.Base(page)
		//save names and parse those files
		ts, err := template.New(fileName).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[fileName] = ts
	}

	return myCache, nil
}

// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	//check to see if we alredy have the template in our cache
// 	_, inMap := tc[t]
// 	if !inMap {
// 		// need to create the template
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		//we have the template in the chache
// 		log.Println("using cached template")
// 	}

// 	tmpl = tc[t]

// 	err = tmpl.Execute(w, nil)
// }

// func createTemplateCache(t string) error {
// 	log.Println("creating template and adding to cache")
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.html",
// 	}

// 	//parse the template
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	// add template to cache (map)
// 	tc[t] = tmpl

// 	return nil
// }
