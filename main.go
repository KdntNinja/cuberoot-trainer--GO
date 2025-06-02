package main

import (
    "fmt"
    "net/http"
    "os"

    "github.com/a-h/templ"
    "github.com/axzilla/templui-quickstart/assets"
    "github.com/axzilla/templui-quickstart/ui/pages"
)

func main() {
    mux := http.NewServeMux()
    SetupAssetsRoutes(mux)
    mux.Handle("GET /", templ.Handler(pages.Landing()))
    mux.Handle("GET /learn", templ.Handler(pages.Learn()))
    mux.HandleFunc("GET /practice", pages.HandlePractice)
    mux.HandleFunc("POST /practice", pages.HandlePractice)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8090"
    }
    fmt.Println("Server is running on http://localhost:" + port)
    err := http.ListenAndServe(":" + port, mux)
    if err != nil {
        fmt.Printf("Server failed to start: %v\n", err)
        os.Exit(1)
    }
}
