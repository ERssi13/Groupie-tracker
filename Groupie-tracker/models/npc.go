package models

type Quest struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Reward      string `json:"reward"`
    Status      string `json:"status"`
}

type NPC struct {
    ID          string   `json:"id"`
    Name        string   `json:"name"`
    Image       string   `json:"image"`
    Description string   `json:"description"`
    Location    string   `json:"location"`
    Role        string   `json:"role"`
    Quotes      []string `json:"quotes"`
    Quests      []Quest  `json:"quests"`
    IsHostile   bool     `json:"isHostile"`
    IsMerchant  bool     `json:"isMerchant"`
}

type NPCResponse struct {
    Success bool  `json:"success"`
    Count   int   `json:"count"`
    Total   int   `json:"total"`
    Data    []NPC `json:"data"`
}

type NPCFilters struct {
    Role     string
    Location string
    Merchant bool
    Hostile  bool
}

type NPCListData struct {
    NPCs        []NPC
    CurrentPage int
    TotalPages  int
    HasNext     bool
    HasPrev     bool
    SearchQuery string
    Filters     NPCFilters
    Title       string
}

type NPCDetailData struct {
    NPC     NPC
    Title   string
    Related []NPC
}