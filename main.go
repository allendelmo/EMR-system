package main

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))
var ddl string

func dashboard(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "dashboard.html", nil)

}
func initDB() {
	ctx := context.Background()
	var err error

	DB, err := sql.Open("sqlite3", "./DB.db")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := DB.ExecContext(ctx, ddl); err != nil {
		fmt.Printf("Error creating table: %q: %s\n", err, ddl)
	}

}

func main() {
	const port = "8080"
	mux := http.NewServeMux()

	mux.HandleFunc("/", dashboard)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Println("Server started at :", port)
	log.Fatal(server.ListenAndServe())
}
