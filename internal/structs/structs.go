package structs

type Pokedex map[string]PokemonS

type Entry struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationS struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous string  `json:"previous"`
	Entrys   []Entry `json:"Entrys"`
}

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
