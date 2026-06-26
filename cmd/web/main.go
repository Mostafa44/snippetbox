package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// Define an application struct to hold the application-wide dependencies for the web application.

type application struct {
	logger *slog.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home) // restrict this route to exact matches on / only
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)

	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	logger.Info("Starting server", "addr", *addr)
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
