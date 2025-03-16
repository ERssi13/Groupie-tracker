package main

import (
    "log"
    "net/http"
    "os"
    "time"
    "Groupe/handlers"
)

func main() {
    r := http.NewServeMux()
    config := handlers.NewConfig()
    fileServer := http.FileServer(http.Dir("static"))
    r.Handle("/static/", http.StripPrefix("/static/", fileServer))
    r.HandleFunc("/", config.LoggerMiddleware(config.HandleHome))
    r.HandleFunc("/bosses", config.LoggerMiddleware(config.HandleBossList))
    r.HandleFunc("/bosses/", config.LoggerMiddleware(config.HandleBossDetail))
    r.HandleFunc("/npcs", config.LoggerMiddleware(config.HandleNPCList))
    r.HandleFunc("/npcs/", config.LoggerMiddleware(config.HandleNPCDetail))
    r.HandleFunc("/api/favorites", config.LoggerMiddleware(config.HandleFavorites))
    srv := &http.Server{
        Handler:      r,
        Addr:         getPort(),
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
        IdleTimeout:  60 * time.Second,
    }
    log.Printf("Server starting on %s", srv.Addr)
    if err := srv.ListenAndServe(); err != nil {
        log.Fatal(err)
    }
}
func getPort() string {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    return ":" + port
}