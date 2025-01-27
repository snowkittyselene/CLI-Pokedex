package commands

import (
	"fmt"
)

func commandMap(c *Config) error {
	res, err := c.Client.CallLocationArea(c.Next)
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
	if c.Previous == nil {
		fmt.Println("You're on the first page")
	}
	res, err := c.Client.CallLocationArea(c.Previous)
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
