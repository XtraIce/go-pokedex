package cli

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	enums "github.com/xtraice/pokedexcli/internal/enums"
)

// getPokeRequest makes a GET request to the PokeAPI and returns the response body as a byte slice.
// It takes a PokeEndpoint and an idName as parameters.
// If the PokeEndpoint is not PokemonLocationArea, it constructs the request URL using the PokeEndpoint and idName.
// If the PokeEndpoint is PokemonLocationArea, it constructs the request URL using the idName.
// It returns the response body as a byte slice and an error if any.
func getPokeRequest(pe enums.PokeEndpoint, idName string) ([]byte, error) {
	var req []byte
	var ok error
	if pe != enums.PokemonLocationArea {
		if req, ok = getRequest(fmt.Sprintf("https://pokeapi.co/api/v2/%s/%s", pe.String(), idName)); ok != nil {
			return nil, ok
		}
	} else {
		if req, ok = getRequest(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/encounters", idName)); ok != nil {
			return nil, ok
		}
	}
	return req, nil
}

// getRequest makes a GET request to the specified URL and returns the response body as a byte slice.
// It takes a request URL as a parameter.
// It returns the response body as a byte slice and an error if any.
func getRequest(request string) ([]byte, error) {
	res, err := http.Get(request)
	if err != nil {
		return nil, errors.New("Error getting request")
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil,
			errors.New(
				fmt.Sprintf(
					"Response failed with status code: %d and\nbody: %s\n",
					res.StatusCode, body))
	}
	if err != nil {
		return nil, errors.New("Error reading response body")
	}
	return body, nil
}
