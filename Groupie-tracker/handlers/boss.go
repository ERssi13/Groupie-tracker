package handlers

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "net/url"
    "strconv"
    "strings"
    "Groupe/models"
)

func (c *Config) HandleBossList(w http.ResponseWriter, r *http.Request) {
    page, _ := strconv.Atoi(r.URL.Query().Get("page"))
    if page < 1 {
        page = 1
    }

    searchQuery := r.URL.Query().Get("name")
    filters := getBossFilters(r)
    
    bosses, total := c.fetchBosses(searchQuery, filters, page)
    
    pageSize := 12
    data := models.BossListData{
        Bosses:      bosses,
        CurrentPage: page,
        TotalPages:  (total + pageSize - 1) / pageSize,
        HasNext:     page*pageSize < total,
        HasPrev:     page > 1,
        SearchQuery: searchQuery,
        Filters:     filters,
        Title:       "Bosses - Elden Ring",
    }

    c.RenderTemplate(w, "bosses", data)
}

func getBossFilters(r *http.Request) models.BossFilters {
    return models.BossFilters{
        Region:   r.URL.Query().Get("region"),
        Category: r.URL.Query().Get("category"),
    }
}

func (c *Config) HandleBossDetail(w http.ResponseWriter, r *http.Request) {
    id := strings.TrimPrefix(r.URL.Path, "/bosses/")
    
    boss, err := c.fetchBossById(id)
    if err != nil {
        http.NotFound(w, r)
        return
    }

    relatedBosses, _ := c.fetchRelatedBosses(boss)
    
    data := models.BossDetailData{
        Boss:    boss,
        Related: relatedBosses,
        Title:   fmt.Sprintf("%s - Elden Ring", boss.Name),
    }

    c.RenderTemplate(w, "boss_detail", data)
}

func (c *Config) fetchBosses(search string, filters models.BossFilters, page int) ([]models.Boss, int) {
    query := url.Values{}
    query.Add("limit", "12")
    query.Add("page", strconv.Itoa(page))
    
    if search != "" {
        query.Add("name", search)
    }
    
    if filters.Region != "" {
        query.Add("region", filters.Region)
    }
    
    if filters.Category != "" {
        query.Add("category", filters.Category)
    }

    url := fmt.Sprintf("%s/bosses?%s", c.BaseURL, query.Encode())
    resp, err := http.Get(url)
    if err != nil {
        log.Printf("Error fetching bosses: %v", err)
        return nil, 0
    }
    defer resp.Body.Close()

    var result models.BossResponse
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        log.Printf("Error decoding boss response: %v", err)
        return nil, 0
    }

    return result.Data, result.Total
}

func (c *Config) fetchBossById(id string) (models.Boss, error) {
    url := fmt.Sprintf("%s/bosses/%s", c.BaseURL, id)
    resp, err := http.Get(url)
    if err != nil {
        return models.Boss{}, err
    }
    defer resp.Body.Close()

    var result struct {
        Data models.Boss `json:"data"`
    }
    
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return models.Boss{}, err
    }

    return result.Data, nil
}

func (c *Config) fetchRelatedBosses(boss models.Boss) ([]models.Boss, error) {
    query := url.Values{}
    query.Add("region", boss.Region)
    query.Add("limit", "3")

    url := fmt.Sprintf("%s/bosses?%s", c.BaseURL, query.Encode())
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result models.BossResponse
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    return result.Data, nil
}