module github.com/snowkittyselene/commands

go 1.23.5

require (
    github.com/snowkittyselene/pokeapi v0.0.0
)

replace (
    github.com/snowkittyselene/pokeapi v0.0.0 => ../pokeapi
)
