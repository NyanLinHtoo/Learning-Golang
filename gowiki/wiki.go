package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

// Function to treat content as safe HTML
func safeHTML(content []byte) template.HTML {
	return template.HTML(content) // Marks content as safe HTML
}

// Use html/template package for rendering HTML templates
var templates = template.Must(template.New("").Funcs(template.FuncMap{
	"safeHTML": safeHTML,
}).ParseFiles("./tmpl/edit.html", "./tmpl/view.html"))

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile("./data/"+filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile("./data/" + filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func createWikiPageLink(p *Page) {
	// Regular expression to match [PageName]
	wikiLinkRegex := regexp.MustCompile(`\[([a-zA-Z0-9]+)\]`)

	// Replace each [PageName] with an <a> tag linking to the corresponding page
	p.Body = wikiLinkRegex.ReplaceAllFunc(p.Body, func(match []byte) []byte {
		// Extract the page name (e.g., "PageName" from "[PageName]")
		pageName := match[1 : len(match)-1]
		// Create the replacement link: <a href="/view/PageName">PageName</a>
		link := []byte(`<a href="/view/` + string(pageName) + `">` + string(pageName) + `</a>`)
		return link
	})
}

func makeHandle(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		// If in "/view/" page is not found, redirect to existing page to edit route
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	createWikiPageLink(p)
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	// FormValue calls Request.ParseMultipartForm and Request.ParseForm if necessary
	body := r.FormValue("body")

	// fmt.Printf("Received body content:\n%s\n", body) // Debug print
	p := &Page{Title: title, Body: []byte(body)}
	createWikiPageLink(p)
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// Convert the Body to a string before passing it to the template
	data := struct {
		Title string
		Body  string
	}{
		Title: p.Title,
		Body:  string(p.Body),
	}
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
