package utils

import (
    "encoding/json"
    "os"
)

type Favorites struct {
    Bosses []string `json:"bosses"`
    NPCs   []string `json:"npcs"`
}

func LoadFavorites() (*Favorites, error) {
    favorites := &Favorites{
        Bosses: make([]string, 0),
        NPCs:   make([]string, 0),
    }

    data, err := os.ReadFile("favorites.json")
    if err != nil {
        if os.IsNotExist(err) {
            return favorites, nil
        }
        return nil, err
    }

    err = json.Unmarshal(data, favorites)
    if err != nil {
        return nil, err
    }

    return favorites, nil
}

func SaveFavorites(favorites *Favorites) error {
    data, err := json.Marshal(favorites)
    if err != nil {
        return err
    }

    return os.WriteFile("favorites.json", data, 0644)
}