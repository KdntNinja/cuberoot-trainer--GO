package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/KdntNinja/cuberoot-trainer--GO/assets"
	"github.com/KdntNinja/cuberoot-trainer--GO/ui/pages"
	"github.com/a-h/templ"
)

func main() {
	mux := http.NewServeMux()
	SetupAssetsRoutes(mux)
	mux.Handle("GET /", templ.Handler(pages.Landing()))
	mux.Handle("GET /learn", templ.Handler(pages.Learn()))

	// Practice page - handle both GET and POST requests separately
	mux.HandleFunc("GET /practice", pages.HandlePractice)
	mux.HandleFunc("POST /practice", pages.HandlePractice)
	fmt.Println("Server is running on http://localhost:8090")
	err := http.ListenAndServe(":8090", mux)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
		os.Exit(1)
	}
}

func SetupAssetsRoutes(mux *http.ServeMux) {
	var isDevelopment = os.Getenv("GO_ENV") != "production"

	assetHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isDevelopment {
			w.Header().Set("Cache-Control", "no-store")
		}

		var fs http.Handler
		if isDevelopment {
			fs = http.FileServer(http.Dir("./assets"))
		} else {
			fs = http.FileServer(http.FS(assets.Assets))
		}

		fs.ServeHTTP(w, r)
	})

	mux.Handle("GET /assets/", http.StripPrefix("/assets/", assetHandler))
}
