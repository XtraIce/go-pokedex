package cli

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	enums "github.com/xtraice/pokedexcli/internal/enums"
)

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
