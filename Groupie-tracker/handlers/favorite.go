package handlers

import (
    "encoding/json"
    "net/http"
    "Groupe/utils"
)

func (c *Config) HandleFavorites(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodPost:
        c.handleAddFavorite(w, r)
    case http.MethodGet:
        c.handleGetFavorites(w, r)
    case http.MethodDelete:
        c.handleRemoveFavorite(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func (c *Config) handleAddFavorite(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Type string `json:"type"`
        ID   string `json:"id"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    favorites, err := utils.LoadFavorites()
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    switch req.Type {
    case "boss":
        favorites.Bosses = append(favorites.Bosses, req.ID)
    case "npc":
        favorites.NPCs = append(favorites.NPCs, req.ID)
    default:
        http.Error(w, "Invalid type", http.StatusBadRequest)
        return
    }

    if err := utils.SaveFavorites(favorites); err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func (c *Config) handleGetFavorites(w http.ResponseWriter, r *http.Request) {
    favorites, err := utils.LoadFavorites()
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    c.JsonResponse(w, http.StatusOK, favorites)
}

func (c *Config) handleRemoveFavorite(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Type string `json:"type"`
        ID   string `json:"id"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    favorites, err := utils.LoadFavorites()
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    switch req.Type {
    case "boss":
        favorites.Bosses = removeFromSlice(favorites.Bosses, req.ID)
    case "npc":
        favorites.NPCs = removeFromSlice(favorites.NPCs, req.ID)
    default:
        http.Error(w, "Invalid type", http.StatusBadRequest)
        return
    }

    if err := utils.SaveFavorites(favorites); err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func removeFromSlice(slice []string, element string) []string {
    for i, v := range slice {
        if v == element {
            return append(slice[:i], slice[i+1:]...)
        }
    }
    return slice
}