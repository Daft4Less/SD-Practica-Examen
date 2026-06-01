package main
import (
    "fmt"
    "net/http"
)
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Servidor Go minimalista (Multi-stage)")
    })
    http.ListenAndServe(":8080", nil)
}