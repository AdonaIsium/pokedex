package pokeapi

type PokemonResp struct {
	ID                     int                `json:"id"`
	Name                   string             `json:"name"`
	BaseExperience         int                `json:"base_experience"`
	Height                 int                `json:"height"`
	IsDefault              bool               `json:"is_default"`
	Order                  int                `json:"order"`
	Weight                 int                `json:"weight"`
	Abilities              []Ability          `json:"abilities"`
	Forms                  []NamedAPIResource `json:"forms"`
	GameIndices            []GameIndex        `json:"game_indices"`
	HeldItems              []HeldItem         `json:"held_items"`
	LocationAreaEncounters string             `json:"location_area_encounters"`
	Moves                  []Move             `json:"moves"`
	Species                NamedAPIResource   `json:"species"`
	Sprites                Sprites            `json:"sprites"`
	Cries                  Cries              `json:"cries"`
	Stats                  []Stat             `json:"stats"`
	Types                  []TypeSlot         `json:"types"`
	PastTypes              []PastType         `json:"past_types"`
	PastAbilities          []PastAbility      `json:"past_abilities"`
}

type Ability struct {
	IsHidden bool             `json:"is_hidden"`
	Slot     int              `json:"slot"`
	Ability  NamedAPIResource `json:"ability"`
}

type GameIndex struct {
	GameIndex int              `json:"game_index"`
	Version   NamedAPIResource `json:"version"`
}

type HeldItem struct {
	Item           NamedAPIResource    `json:"item"`
	VersionDetails []ItemVersionDetail `json:"version_details"`
}

type ItemVersionDetail struct {
	Rarity  int              `json:"rarity"`
	Version NamedAPIResource `json:"version"`
}

type Move struct {
	Move                NamedAPIResource    `json:"move"`
	VersionGroupDetails []MoveVersionDetail `json:"version_group_details"`
}

type MoveVersionDetail struct {
	LevelLearnedAt  int              `json:"level_learned_at"`
	VersionGroup    NamedAPIResource `json:"version_group"`
	MoveLearnMethod NamedAPIResource `json:"move_learn_method"`
	Order           int              `json:"order"`
}

type Sprites struct {
	BackDefault      string          `json:"back_default"`
	BackFemale       *string         `json:"back_female"`
	BackShiny        string          `json:"back_shiny"`
	BackShinyFemale  *string         `json:"back_shiny_female"`
	FrontDefault     string          `json:"front_default"`
	FrontFemale      *string         `json:"front_female"`
	FrontShiny       string          `json:"front_shiny"`
	FrontShinyFemale *string         `json:"front_shiny_female"`
	Other            OtherSprites    `json:"other"`
	Versions         VersionsSprites `json:"versions"`
}

type OtherSprites struct {
	DreamWorld      VersionSpritesEntry `json:"dream_world"`
	Home            VersionSpritesEntry `json:"home"`
	OfficialArtwork VersionSpritesEntry `json:"official-artwork"`
	Showdown        VersionSpritesEntry `json:"showdown"`
}

type VersionsSprites struct {
	GenerationI    GenerationI    `json:"generation-i"`
	GenerationII   GenerationII   `json:"generation-ii"`
	GenerationIII  GenerationIII  `json:"generation-iii"`
	GenerationIV   GenerationIV   `json:"generation-iv"`
	GenerationV    GenerationV    `json:"generation-v"`
	GenerationVI   GenerationVI   `json:"generation-vi"`
	GenerationVII  GenerationVII  `json:"generation-vii"`
	GenerationVIII GenerationVIII `json:"generation-viii"`
}

type GenerationI struct {
	RedBlue VersionSpritesEntry `json:"red-blue"`
	Yellow  VersionSpritesEntry `json:"yellow"`
}

type GenerationII struct {
	Crystal VersionSpritesEntry `json:"crystal"`
	Gold    VersionSpritesEntry `json:"gold"`
	Silver  VersionSpritesEntry `json:"silver"`
}

type GenerationIII struct {
	Emerald          VersionSpritesEntry `json:"emerald"`
	FireredLeafgreen VersionSpritesEntry `json:"firered-leafgreen"`
	RubySapphire     VersionSpritesEntry `json:"ruby-sapphire"`
}

type GenerationIV struct {
	DiamondPearl        VersionSpritesEntry `json:"diamond-pearl"`
	HeartgoldSoulsilver VersionSpritesEntry `json:"heartgold-soulsilver"`
	Platinum            VersionSpritesEntry `json:"platinum"`
}

type GenerationV struct {
	BlackWhite BlackWhiteSprites `json:"black-white"`
}
type GenerationVI struct {
	OmegarubyAlphasapphire VersionSpritesEntry `json:"omegaruby-alphasapphire"`
	XY                     VersionSpritesEntry `json:"x-y"`
}
type GenerationVII struct {
	Icons             VersionSpritesEntry `json:"icons"`
	UltraSunUltraMoon VersionSpritesEntry `json:"ultra-sun-ultra-moon"`
}
type GenerationVIII struct {
	Icons VersionSpritesEntry `json:"icons"`
}

type VersionSpritesEntry struct {
	BackDefault      string  `json:"back_default,omitempty"`
	BackFemale       *string `json:"back_female,omitempty"`
	BackShiny        string  `json:"back_shiny,omitempty"`
	BackShinyFemale  *string `json:"back_shiny_female,omitempty"`
	FrontDefault     string  `json:"front_default,omitempty"`
	FrontFemale      *string `json:"front_female,omitempty"`
	FrontShiny       string  `json:"front_shiny,omitempty"`
	FrontShinyFemale *string `json:"front_shiny_female,omitempty"`
}

type BlackWhiteSprites struct {
	Animated         *VersionSpritesEntry `json:"animated,omitempty"`
	BackDefault      string               `json:"back_default"`
	BackFemale       *string              `json:"back_female"`
	BackShiny        string               `json:"back_shiny"`
	BackShinyFemale  *string              `json:"back_shiny_female"`
	FrontDefault     string               `json:"front_default"`
	FrontFemale      *string              `json:"front_female"`
	FrontShiny       string               `json:"front_shiny"`
	FrontShinyFemale *string              `json:"front_shiny_female"`
}

type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type Stat struct {
	BaseStat int              `json:"base_stat"`
	Effort   int              `json:"effort"`
	Stat     NamedAPIResource `json:"stat"`
}

type TypeSlot struct {
	Slot int              `json:"slot"`
	Type NamedAPIResource `json:"type"`
}

type PastType struct {
	Generation NamedAPIResource `json:"generation"`
	Types      []TypeSlot       `json:"types"`
}

type PastAbility struct {
	Generation NamedAPIResource   `json:"generation"`
	Abilities  []PastAbilityEntry `json:"abilities"`
}

type PastAbilityEntry struct {
	Ability  *NamedAPIResource `json:"ability"`
	IsHidden bool              `json:"is_hidden"`
	Slot     int               `json:"slot"`
}
