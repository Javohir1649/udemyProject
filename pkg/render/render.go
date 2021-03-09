package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

// RenderTemplate html fayllarni parse qiladi
func RenderTemplate(w http.ResponseWriter, html string) {

	_, err := RenderTemplateTest(w)
	if err != nil {
		fmt.Println("error getting template cache")
	}

	parsedTemplate, _ := template.ParseFiles("../../templates/" + html)
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}

func RenderTemplateTest(w http.ResponseWriter) (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("../../templates/*.page.html")
	fmt.Printf("%v", pages)
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {

		name := filepath.Base(page)
		fmt.Printf("%v", name)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		fmt.Printf("%v", ts)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("../../templates/*.layout.html")
		fmt.Printf("%v", matches)
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../templates/*.layout.html")
			fmt.Printf("%v", ts)
			if err != nil {
				return myCache, err
			}

		}

		myCache[name] = ts
	}

	return myCache, nil

}
