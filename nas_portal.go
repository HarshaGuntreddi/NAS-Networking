package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/gorilla/mux"
)

type Service struct {
    Name string `json:"name"`
    Url  string `json:"url"`
    Up   bool   `json:"up"`
}

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        _ = json.NewEncoder(w).Encode(map[string]string{
            "status": "ok",
            "ts":     time.Now().UTC().Format(time.RFC3339),
        })
    }).Methods("GET")

    r.HandleFunc("/api/services", func(w http.ResponseWriter, r *http.Request) {
        // stub: pretend to probe truenas, plex, etc
        services := []Service{
            {"truenas", "http://truenas1", true},
            {"plex", "http://nas.local/plex", true},
            {"grafana", "http://nas.local:3000", true},
        }
        w.Header().Set("Content-Type", "application/json")
        _ = json.NewEncoder(w).Encode(services)
    }).Methods("GET")

    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./nas_static")))

    addr := ":8080"
    if v := os.Getenv("PORT"); v != "" {
        addr = ":" + v
    }
    log.Printf("nas portal listening on %s", addr)
    log.Fatal(http.ListenAndServe(addr, r))
}
