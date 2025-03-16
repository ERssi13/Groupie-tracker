package models

type Boss struct {
    ID          string   `json:"id"`
    Name        string   `json:"name"`
    Image       string   `json:"image"`
    Description string   `json:"description"`
    Location    string   `json:"location"`
    Region      string   `json:"region"`
    Drops       []string `json:"drops"`
    HealthPoints int     `json:"healthPoints"`
    Category    string   `json:"category"`
    IsDefeated  bool    `json:"isDefeated"`
}

type BossResponse struct {
    Success bool   `json:"success"`
    Count   int    `json:"count"`
    Total   int    `json:"total"`
    Data    []Boss `json:"data"`
}

type BossFilters struct {
    Region   string
    Category string
    Defeated string
}

type BossListData struct {
    Bosses      []Boss
    CurrentPage int
    TotalPages  int
    HasNext     bool
    HasPrev     bool
    SearchQuery string
    Filters     BossFilters
    Title       string
}

type BossDetailData struct {
    Boss  Boss
    Title string
    Related []Boss
}