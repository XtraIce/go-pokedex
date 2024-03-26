package cli

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"sort"
	"time"

	"github.com/xtraice/pokedexcli/internal/enums"
	"github.com/xtraice/pokedexcli/internal/lcache"
	"github.com/xtraice/pokedexcli/internal/structs"
)

var currentPokemon *structs.PokemonS
var pokemons *lcache.Cache
var pokedex structs.Pokedex

// getPokemon retrieves the information of a specific Pokemon.
// It first checks if the Pokemon is already cached in the 'pokemons' cache.
// If it is, it unmarshals the cached data into the 'currentPokemon' variable.
// If it is not, it makes a request to get the Pokemon data and caches it.
// Returns true if the Pokemon information is successfully retrieved, false otherwise.
func getPokemon(pokemon string) bool {
	if currentPokemon == nil {
		pokemons = lcache.NewCache(time.Minute * 10)
	}
	if req, ok := pokemons.Get(pokemon); ok {
		if err := json.Unmarshal(req, &currentPokemon); err != nil {
			return false
		}
	} else if req, err := getPokeRequest(enums.Pokemon, pokemon); err == nil {
		if err := json.Unmarshal(req, &currentPokemon); err != nil {
			return false
		}
		pokemons.Add(pokemon, req)
	}
	return true
}

// getTrainerLevel calculates the total trainer level based on the base experience of each caught Pokemon.
// Returns the trainer level as a float64.
func getTrainerLevel() float64 {
	trainerXp := 0.0
	for _, stats := range pokedex {
		trainerXp += float64(stats.BaseExperience) / 100
	}
	return trainerXp
}

// PokedexList prints the list of caught Pokemon in the Pokedex.
// If no Pokemon is caught yet, it prints a message indicating that.
func PokedexList() {
	if pokedex == nil {
		log.Default().Println("No pokemon caught yet")
		return
	}
	fmt.Println("Your Pokedex:")
	pokeList := []string{}
	for name := range pokedex {
		pokeList = append(pokeList, name)
	}
	sort.Strings(pokeList)
	for _, name := range pokeList {
		fmt.Println(" -", name, " #", pokedex[name].ID)
	}
}

// PokemonCatch attempts to catch a Pokemon.
// It retrieves the Pokemon information using the 'getPokemon' function.
// It calculates the capture rate based on the Pokemon's base experience and the trainer's level.
// It generates a random number between 0 and 100 as the capture difficulty.
// It simulates the catching process by printing dots and waiting for a second.
// If the capture rate is higher than the difficulty, the Pokemon is caught and added to the Pokedex.
// If not, the Pokemon escapes.
func PokemonCatch(pokemon string) {
	if !getPokemon(pokemon) {
		log.Fatal("Error getting pokemon")
	}
	captureRate := 1000.0 / float64(currentPokemon.BaseExperience)
	captureRate += getTrainerLevel()

	// generate random number between 0 and 100
	difficulty := rand.Float32() * 100

	fmt.Print("Trying to catch " + pokemon + " with a pokeball")
	for i := 0; i < 3; i++ {
		fmt.Print(".")
		time.Sleep(time.Second)
	}

	if captureRate > float64(difficulty) {
		fmt.Println(pokemon + " was caught!")
		if pokedex == nil {
			pokedex = map[string]structs.PokemonS{}
		}
		pokedex[pokemon] = *currentPokemon
		fmt.Printf("View %s with the 'inspect' command\n", pokemon)
	} else {
		fmt.Println(pokemon + " escaped!")
		//log.Println("catch rate: ", captureRate)
		//log.Println("random number: ", difficulty)
	}
}

// PokemonInspect displays detailed information about a caught Pokemon.
// It checks if the Pokedex is empty or if the specified Pokemon is not caught yet.
// If the Pokemon is found, it prints its name, height, weight, stats, and types.
func PokemonInspect(pokemon string) {
	if pokedex == nil {
		log.Default().Println("No pokemon caught yet")
		return
	}
	if _, info := pokedex[pokemon]; !info {
		log.Default().Println("Pokemon not caught yet")
		return
	} else {
		fmt.Println("Name: ", pokedex[pokemon].Name)
		fmt.Println("Height: ", pokedex[pokemon].Height)
		fmt.Println("Weight: ", pokedex[pokemon].Weight)
		fmt.Println("Stats:")
		for _, stat := range pokedex[pokemon].Stats {
			fmt.Println(" -", stat.Stat.Name, ": ", stat.BaseStat)
		}
		fmt.Println("Type: ")
		for _, typeInfo := range pokedex[pokemon].Types {
			fmt.Println(" - ", typeInfo.Type.Name)
		}
	}
}
