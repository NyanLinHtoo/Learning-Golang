package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"text/template"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	fmt.Println("Saving page:", p.Title)                                // Debug print
	fmt.Printf("Page content before processing:\n%s\n", string(p.Body)) // Debug print
	createWikiPageLink(p)
	fmt.Printf("Page content after processing:\n%s\n", string(p.Body)) // Debug print

	return os.WriteFile("./data/"+filename, p.Body, 0600)
}

func createWikiPageLink(p *Page) {
	fmt.Println("Processing page:", p.Title)
	var linkPattern = regexp.MustCompile(`\[([a-zA-Z0-9]+)\]`)
	fmt.Printf("Found : %v \n", linkPattern) // Debug print
	bodyString := string(p.Body)
	fmt.Printf("Found : %v \n", bodyString) // Debug print
	matches := linkPattern.FindAll([]byte(bodyString), -1)
	fmt.Printf("Found %d matches\n", len(matches)) // Debug print

	p.Body = linkPattern.ReplaceAllFunc(p.Body, func(match []byte) []byte {
		group := linkPattern.ReplaceAllString(string(match), "$1")
		fmt.Println("Found group:", string(group))
		return []byte(fmt.Sprintf("<a href='/view/%s'>%s</a>", group, group))
	})
	fmt.Println("Finished processing page:", p.Title)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile("./data/" + filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
// 	m := validPath.FindStringSubmatch(r.URL.Path)
// 	if m == nil {
// 		http.NotFound(w, r)
// 		return "", errors.New("invalid Page Title")
// 	}
// 	fmt.Println("m : ", m)
// 	return m[2], nil // The title is the second subexpression.
// }

func makeHandle(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Here we will extract the page title from the Request,
		// and call the provided handler 'fn'

		m := validPath.FindStringSubmatch(r.URL.Path)

		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/view/"):]
	// title, err := getTitle(w, r)
	// if err != nil {
	// 	return
	// }
	p, err := loadPage(title)
	if err != nil {
		// if in "/view/" page is not found, redirect to existing page to edit route
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/edit/"):]
	// title, err := getTitle(w, r)
	// if err != nil {
	// 	return
	// }
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/save/"):]
	// title, err := getTitle(w, r)
	// if err != nil {
	// 	return
	// }
	//  FormValue calls Request.ParseMultipartForm and Request.ParseForm if necessary and ignores any errors returned by these functions. If key is not present, FormValue returns the empty string. To access multiple values of the same key, call ParseForm and then inspect [Request.Form] directly.
	body := r.FormValue("body")

	fmt.Printf("Received body content:\n%s\n", body) // Debug print
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var templates = template.Must(template.ParseFiles("./tmpl/edit.html", "./tmpl/view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// t, err := template.ParseFiles(tmpl + ".html")
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// err = t.Execute(w, p)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/FrontPage", http.StatusFound)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/view/", makeHandle(viewHandler))
	http.HandleFunc("/edit/", makeHandle(editHandler))
	http.HandleFunc("/save/", makeHandle(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
