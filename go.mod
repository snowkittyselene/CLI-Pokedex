module github.com/snowkittyselene/pokedex-cli

go 1.23.5

replace (
	github.com/snowkittyselene/commands v0.0.0 => ./internal/commands
	github.com/snowkittyselene/pokeapi v0.0.0 => ./internal/pokeapi
	github.com/snowkittyselene/pokecache v0.0.0 => ./internal/pokecache
)

require (
	github.com/snowkittyselene/commands v0.0.0
	github.com/snowkittyselene/pokeapi v0.0.0
)

require github.com/snowkittyselene/pokecache v0.0.0 // indirect
