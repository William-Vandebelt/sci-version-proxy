package main

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func getTemplate(fm template.FuncMap, name string) *template.Template {
	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("templates/"+name+".tmpl.html",
		"templates/header.tmpl.html", "templates/footer.tmpl.html"))
	log.Printf("%s template returned\n", name)
	return tpl
}

// HandleError : Func that sends back a 500 error code if there is an error
func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
}

func routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tpl := getTemplate(template.FuncMap{}, "index")
		err := tpl.ExecuteTemplate(w, "index.tmpl.html", "Hello Message From Index.go")
		handleError(w, err)
	})

	return router
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %q\n", err)
	}
}

// Build with >> go build -o bin/sci-version-proxy -v .
func main() {
	port := os.Getenv("PORT")

	r := routes()

	log.Printf("Golang App running...\n")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
