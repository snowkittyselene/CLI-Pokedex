module github.com/snowkittyselene/pokedex-cli

go 1.23.5

replace (
    github.com/snowkittyselene/commands v0.0.0 => ./commands
    github.com/snowkittyselene/poke_api v0.0.0 => ./poke_api

)
require (
    github.com/snowkittyselene/commands v0.0.0
    github.com/snowkittyselene/poke_api v0.0.0 // indirect
)

