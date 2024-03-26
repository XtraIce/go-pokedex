package structs

// Pokedex represents a map of PokemonS structs, where the key is a string.
type Pokedex map[string]PokemonS

// Entry represents a Pokemon entry with its name and URL.
type Entry struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// LocationS represents a location with its count, next and previous URLs, and a list of entries.
type LocationS struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous string  `json:"previous"`
	Entrys   []Entry `json:"Entrys"`
}

// LocationAreaS represents a location area with its ID, name, game index, encounter method rates,
// location, names, and Pokemon encounters.
type LocationAreaS struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod Entry `json:"encounter_method"`
		VersionDetails  []struct {
			Rate    int   `json:"rate"`
			Version Entry `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location Entry `json:"location"`
	Names    []struct {
		Name     string `json:"name"`
		Language Entry  `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon        Entry `json:"pokemon"`
		VersionDetails []struct {
			Version          Entry `json:"version"`
			MaxChance        int   `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          Entry `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
