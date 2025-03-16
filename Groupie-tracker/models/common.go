package models

type PageResponse struct {
    Title       string
    CurrentPage int
    TotalPages  int
    HasNext     bool
    HasPrev     bool
    SearchQuery string
}

type Favorites struct {
    Bosses []string `json:"bosses"`
    NPCs   []string `json:"npcs"`
}

type APIError struct {
    Status  int    `json:"status"`
    Message string `json:"message"`
}

type SearchOptions struct {
    Query  string
    Page   int
    Limit  int
    Filter interface{}
    Sort   string
}