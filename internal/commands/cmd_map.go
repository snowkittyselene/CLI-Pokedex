package commands

import (
	"fmt"

	"github.com/snowkittyselene/pokeapi"
)

const locationArea string = baseApiUrl + "/location-area/"

func commandMap(c *Config) error {
	url := locationArea
	if c.Next != "" {
		url = c.Next
	}
	res, err := pokeapi.CallLocationArea(url)
	if err != nil {
		return err
	}
	c.Next = res.Next
	c.Previous = res.Previous
	for _, location := range res.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapBack(c *Config) error {
	if c.Previous == "" {
		return fmt.Errorf("no previous map page")
	}
	res, err := pokeapi.CallLocationArea(c.Previous)
	if err != nil {
		return err
	}
	c.Next = res.Next
	c.Previous = res.Previous
	for _, location := range res.Results {
		fmt.Println(location.Name)
	}
	return nil
}
