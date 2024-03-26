package cli

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/xtraice/pokedexcli/internal/enums"
	"github.com/xtraice/pokedexcli/internal/lcache"
	"github.com/xtraice/pokedexcli/internal/structs"
)

var currLocations *structs.LocationS
var currlocationArea *structs.LocationAreaS
var currLocationIdx int
var locations *lcache.Cache
var locationAreas *lcache.Cache
var fwd bool

func getLocations() {
	if currLocations == nil { // first time
		locations = lcache.NewCache(time.Minute * 5)
		var req []byte
		var ok error
		if req, ok = getPokeRequest(enums.Location, ""); ok != nil {
			log.Fatal(ok)
		} else if err := json.Unmarshal(req, &currLocations); err != nil {
			log.Fatal(err)
		}
		locations.Add(fmt.Sprint(currLocationIdx), req) // cache
	} else {
		var req []byte
		var ok bool
		var err error
		if req, ok = locations.Get(fmt.Sprint(currLocationIdx)); ok {
			//log.Default().Println("Found in cache")
		} else if req, err = getRequest(currLocations.Next); err == nil { // fetch next
			locations.Add(fmt.Sprint(currLocationIdx), req) // cache
		} else {
			log.Default().Println(err)
		}
		if err = json.Unmarshal(req, &currLocations); err != nil { // unmarshal
			log.Default().Println(err)
		}
	}
	for _, location := range currLocations.Entrys {
		fmt.Println(location.Name)
	}
	fwd = true
	currLocationIdx += len(currLocations.Entrys)
}

func getLocationsBefore() {
	if fwd {
		fwd = false
		currLocationIdx = max(0, currLocationIdx-20)
	}
	currLocationIdx = max(0, currLocationIdx-20)
	if currLocations == nil { // first time or no previous locations
		log.Println("No previous locations")
		return
	}
	if req, ok := locations.Get(fmt.Sprint(currLocationIdx)); ok {
		if err := json.Unmarshal(req, &currLocations); err != nil {
			log.Default().Println(err)
		}
	} else { // shouldn't happen
		log.Fatal("Error fetching previous locations")
		return
	}
	for _, location := range currLocations.Entrys {
		fmt.Println(location.Name)
	}
}

func exploreLocation(locationArea string) {
	if currlocationArea == nil {
		//fmt.Println("First time")
		locationAreas = lcache.NewCache(time.Minute * 5)
	}
	if req, ok := locationAreas.Get(locationArea + "-area"); ok {
		//fmt.Println("Found in cache")
		if err := json.Unmarshal(req, &currlocationArea); err != nil {
			log.Fatal(err)
		}
	} else if req, ok := getPokeRequest(enums.LocationArea, locationArea+"-area"); ok == nil {
		//fmt.Println("Fetching")
		if err := json.Unmarshal(req, &currlocationArea); err != nil {
			log.Fatal(err)
		}
		locationAreas.Add(locationArea+"-area", req)
	}
	fmt.Println("Pokemon encounters in", locationArea)
	for _, encounter := range currlocationArea.PokemonEncounters {
		fmt.Println(" - " + encounter.Pokemon.Name)
	}
}
