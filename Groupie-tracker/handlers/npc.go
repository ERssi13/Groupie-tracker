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

func (c *Config) HandleNPCList(w http.ResponseWriter, r *http.Request) {
    page, _ := strconv.Atoi(r.URL.Query().Get("page"))
    if page < 1 {
        page = 1
    }

    searchQuery := r.URL.Query().Get("name")
    filters := getNPCFilters(r)
    
    npcs, total := c.fetchNPCs(searchQuery, filters, page)
    
    pageSize := 12
    data := models.NPCListData{
        NPCs:        npcs,
        CurrentPage: page,
        TotalPages:  (total + pageSize - 1) / pageSize,
        HasNext:     page*pageSize < total,
        HasPrev:     page > 1,
        SearchQuery: searchQuery,
        Filters:     filters,
        Title:       "NPCs - Elden Ring",
    }

    c.RenderTemplate(w, "npcs", data)
}

func getNPCFilters(r *http.Request) models.NPCFilters {
    return models.NPCFilters{
        Role:     r.URL.Query().Get("role"),
        Location: r.URL.Query().Get("location"),
    }
}

func (c *Config) HandleNPCDetail(w http.ResponseWriter, r *http.Request) {
    id := strings.TrimPrefix(r.URL.Path, "/npcs/")
    
    npc, err := c.fetchNPCById(id)
    if err != nil {
        http.NotFound(w, r)
        return
    }

    relatedNPCs, _ := c.fetchRelatedNPCs(npc)
    
    data := models.NPCDetailData{
        NPC:     npc,
        Related: relatedNPCs,
        Title:   fmt.Sprintf("%s - Elden Ring", npc.Name),
    }

    c.RenderTemplate(w, "npc_detail", data)
}

func (c *Config) fetchNPCs(search string, filters models.NPCFilters, page int) ([]models.NPC, int) {
    query := url.Values{}
    query.Add("limit", "12")
    query.Add("page", strconv.Itoa(page))
    
    if search != "" {
        query.Add("name", search)
    }
    
    if filters.Role != "" {
        query.Add("role", filters.Role)
    }
    
    if filters.Location != "" {
        query.Add("location", filters.Location)
    }

    url := fmt.Sprintf("%s/npcs?%s", c.BaseURL, query.Encode())
    resp, err := http.Get(url)
    if err != nil {
        log.Printf("Error fetching NPCs: %v", err)
        return nil, 0
    }
    defer resp.Body.Close()

    var result models.NPCResponse
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        log.Printf("Error decoding NPC response: %v", err)
        return nil, 0
    }

    return result.Data, result.Total
}

func (c *Config) fetchNPCById(id string) (models.NPC, error) {
    url := fmt.Sprintf("%s/npcs/%s", c.BaseURL, id)
    resp, err := http.Get(url)
    if err != nil {
        return models.NPC{}, err
    }
    defer resp.Body.Close()

    var result struct {
        Data models.NPC `json:"data"`
    }
    
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return models.NPC{}, err
    }

    return result.Data, nil
}

func (c *Config) fetchRelatedNPCs(npc models.NPC) ([]models.NPC, error) {
    query := url.Values{}
    query.Add("role", npc.Role)
    query.Add("limit", "3")

    url := fmt.Sprintf("%s/npcs?%s", c.BaseURL, query.Encode())
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result models.NPCResponse
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    return result.Data, nil
}