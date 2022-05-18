// Fetches board game data from Board Game Atlas.
// See: https://www.boardgameatlas.com/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Game struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	ImageURL    string   `json:"image_url"`
	Year        int64    `json:"year_published"`
	MinPlayers  int64    `json:"min_players"`
	MaxPlayers  int64    `json:"max_players"`
	MinPlayTime int64    `json:"min_playtime"`
	MaxPlayTime int64    `json:"max_playtime"`
	MinAge      int64    `json:"min_age"`
	Description string   `json:"description"`
	Mechanics   []string `json:"mechanics"`
	Categories  []string `json:"categories"`
	Images      struct {
		Thumb    string `json:"thumb"`
		Small    string `json:"small"`
		Medium   string `json:"medium"`
		Large    string `json:"large"`
		Original string `json:"original"`
	} `json:"images"`
}

func buildURL(gameIDs []string, clientID string) string {
	ids := strings.Join(gameIDs, ",")
	return fmt.Sprintf("https://api.boardgameatlas.com/api/search?ids=%s&client_id=%s", ids, clientID)
}

func fetch(url string, dest any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&dest)
	if err != nil {
		return err
	}

	return nil
}

func fetchGames(gameIDs []string, clientID string, categories, mechanics map[string]string) ([]*Game, error) {
	type response struct {
		Games []struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			ImageURL    string `json:"image_url"`
			Year        int64  `json:"year_published"`
			MinPlayers  int64  `json:"min_players"`
			MaxPlayers  int64  `json:"max_players"`
			MinPlayTime int64  `json:"min_playtime"`
			MaxPlayTime int64  `json:"max_playtime"`
			MinAge      int64  `json:"min_age"`
			Description string `json:"description"`
			Images      struct {
				Thumb    string `json:"thumb"`
				Small    string `json:"small"`
				Medium   string `json:"medium"`
				Large    string `json:"large"`
				Original string `json:"original"`
			} `json:"images"`
			Mechanics []struct {
				ID string `json:"id"`
			} `json:"mechanics"`
			Categories []struct {
				ID string `json:"id"`
			} `json:"categories"`
		} `json:"games"`
	}

	ids := strings.Join(gameIDs, ",")

	u := fmt.Sprintf("https://api.boardgameatlas.com/api/search?ids=%s&client_id=%s", ids, clientID)

	var res response
	err := fetch(u, &res)
	if err != nil {
		return nil, err
	}

	games := []*Game{}

	for _, game := range res.Games {
		g := &Game{}
		g.ID = game.ID
		g.Name = game.Name
		g.ImageURL = game.ImageURL
		g.Year = game.Year
		g.MinPlayers = game.MinPlayers
		g.MaxPlayers = game.MaxPlayers
		g.MinPlayTime = game.MinPlayTime
		g.MaxPlayTime = game.MaxPlayTime
		g.MinAge = game.MinAge
		g.Description = game.Description
		g.Images = game.Images

		for _, m := range game.Mechanics {
			if v, ok := mechanics[m.ID]; ok {
				g.Mechanics = append(g.Mechanics, v)
			}
		}

		for _, c := range game.Categories {
			if v, ok := categories[c.ID]; ok {
				g.Categories = append(g.Categories, v)
			}
		}

		games = append(games, g)
	}

	return games, nil
}

func fetchCategories(clientID string) (map[string]string, error) {
	type response struct {
		Categories []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"categories"`
	}

	u := fmt.Sprintf("https://api.boardgameatlas.com/api/game/categories?client_id=%s", clientID)

	var res response
	err := fetch(u, &res)
	if err != nil {
		return nil, err
	}

	categories := map[string]string{}
	for _, category := range res.Categories {
		categories[category.ID] = category.Name
	}

	return categories, nil
}

func fetchMechanics(clientID string) (map[string]string, error) {
	type response struct {
		Categories []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"mechanics"`
	}

	u := fmt.Sprintf("https://api.boardgameatlas.com/api/game/mechanics?client_id=%s", clientID)

	var res response
	err := fetch(u, &res)
	if err != nil {
		return nil, err
	}

	mechanics := map[string]string{}
	for _, mechanic := range res.Categories {
		mechanics[mechanic.ID] = mechanic.Name
	}

	return mechanics, nil
}

func main() {
	clientID := flag.String("cid", "", "Board Game Atlas client id")
	flag.Parse()

	if *clientID == "" {
		flag.Usage()
		os.Exit(0)
	}

	gameIDs := []string{
		"6VQXkkC5ql", // Dominion
		"OIXt3DmJU0", // Catan
		"OF145SrX44", // 7 Wonders
		"H3yXWu5No0", // Dixit
		"fDn9rQjH9O", // Terraforming Mars
		"RLlDWHh7hR", // Gloomhaven
		"yqR4PtpO8X", // Scythe
		"5H5JS0KLzK", // Wingspan
		"Zd3nGqEu4q", // Arkham Horror
		"Zspd2nNHfz", // Mansion of Madness
		"ZwlnfX1J2z", // Eldritch Horror
		"XzBSrq35uz", // Memoir '44
		"CWPEdkpzeb", // Eclipse
		"vFMkVLRu8x", // Ethnos
		"6dSwjFyCwc", // The Grizzled
		"dgZDurgbuY", // Munchkin
		"Pmb6ULttIc", // Western Legends
		"6FmFeux5xH", // Pandemic
		"oGVgRSAKwX", // Carcassonne
		"AuBvbISHR6", // Ticket to Ride
		"ZF9FOLsqys", // The Lord of the Rings: The Card Game
	}

	categories, err := fetchCategories(*clientID)
	if err != nil {
		log.Fatal(err)
	}

	mechanics, err := fetchMechanics(*clientID)
	if err != nil {
		log.Fatal(err)
	}

	games, err := fetchGames(gameIDs, *clientID, categories, mechanics)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.MarshalIndent(map[string]any{"games": games}, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(b)
}
