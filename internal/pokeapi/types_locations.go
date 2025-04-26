package pokeapi

type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// RespLocationDetails represents the full payload for a location-area.
type RespLocationDetails struct {
	ID                   int                   `json:"id"`
	Name                 string                `json:"name"`
	GameIndex            int                   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             NamedAPIResource      `json:"location"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource    `json:"encounter_method"`
	VersionDetails  []VersionDetailRate `json:"version_details"`
}

type VersionDetailRate struct {
	Rate    int              `json:"rate"`
	Version NamedAPIResource `json:"version"`
}

type Name struct {
	Name     string           `json:"name"`
	Language NamedAPIResource `json:"language"`
}

type PokemonEncounter struct {
	Pokemon        NamedAPIResource       `json:"pokemon"`
	VersionDetails []PokemonVersionDetail `json:"version_details"`
}

type PokemonVersionDetail struct {
	Version          NamedAPIResource   `json:"version"`
	MaxChance        int                `json:"max_chance"`
	ConditionValues  []NamedAPIResource `json:"condition_values"`
	EncounterDetails []EncounterDetail  `json:"encounter_details"`
}

type EncounterDetail struct {
	MinLevel int              `json:"min_level"`
	MaxLevel int              `json:"max_level"`
	Chance   int              `json:"chance"`
	Method   NamedAPIResource `json:"method"`
}
