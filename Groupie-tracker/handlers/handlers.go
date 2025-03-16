package handlers

import (
    "html/template"
    "net/http"
    "encoding/json"
    "log"
    "path/filepath"
)

// Fonction pour soustraire deux nombres
func sub(a, b int) int {
    return a - b
}

// Fonction pour additionner deux nombres
func add(a, b int) int {
    return a + b
}

type Config struct {
    Templates *template.Template
    BaseURL   string
}

func NewConfig() *Config {
    tmpl := template.Must(template.New("").Funcs(template.FuncMap{
        "sub": sub, 
        "add": add,  
    }).ParseGlob(filepath.Join("templates", "*.html")))

    return &Config{
        Templates: tmpl,
        BaseURL:   "https://eldenring.fanapis.com/api",
    }
}

func (c *Config) RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    err := c.Templates.ExecuteTemplate(w, tmpl, data)
    if err != nil {
        log.Printf("Error rendering template: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

func (c *Config) JsonResponse(w http.ResponseWriter, status int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}

func (c *Config) LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.URL.Path)
        next(w, r)
    }
}
func (c *Config) HandleHome(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    data := struct {
        Title string
    }{
        Title: "Welcome to Elden Ring Wiki",
    }

    c.RenderTemplate(w, "layout.html", data)
}
