package enums

type PokeEndpoint int

const (
	Berries PokeEndpoint = iota + 1
	BerryFirmness
	BerryFlavor
	Contest
	ContestEffect
	ContestSuperEffect
	EncounterMethod
	EncounterCondition
	EncounterConditionValue
	EvolutionChain
	EvolutionTrigger
	GamesGeneration
	GamesPokedex
	GamesVersion
	GamesVersionGroup
	Item
	ItemAttribute
	ItemCategory
	ItemFlingEffect
	ItemPocket
	Location
	LocationArea
	LocationPalParkArea
	LocationRegion
	Machine
	Move
	MoveAilment
	MoveBattleStyle
	MoveCategory
	MoveDamageClass
	MoveLearnMethod
	MoveTarget
	Pokemon
	PokemonAbility
	PokemonCharacteristic
	PokemonEggGroup
	PokemonGender
	PokemonGrowthRate
	PokemonNature
	PokeathlonStat
	PokemonLocationArea
	PokemonColor
	PokemonForm
	PokemonHabitat
	PokemonShape
	PokemonSpecies
	Stat
	Type
	Language
)

func (p PokeEndpoint) String() string {
	return [...]string{"berries", "berry-firmness", "berry-flavor", "contest", "contest-effect",
		"super-contest-effect", "encounter-method", "encounter-condition", "encounter-condition-value",
		"evolution-chain", "evolution-trigger", "generation", "pokedex", "version", "version-group",
		"item", "item-attribute", "item-category", "item-fling-effect", "item-pocket", "location",
		"location-area", "pal-park-area", "region", "machine", "move", "move-ailment", "move-battle-style",
		"move-category", "move-damage-class", "move-learn-method", "move-target", "pokemon", "ability",
		"characteristic", "egg-group", "gender", "growth-rate", "nature", "pokeathlon-stat",
		"pokemon/%s/encounters", "pokemon-color", "pokemon-form", "pokemon-habitat", "pokemon-shape",
		"pokemon-species", "stat", "type", "language"}[p-1]
}

func (p PokeEndpoint) EnumIndex() int {
	return int(p)
}
