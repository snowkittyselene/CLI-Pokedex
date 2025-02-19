package commands

import (
	"fmt"
)

func commandMapForward(c *Config, args ...string) error {
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

func commandMapBack(c *Config, args ...string) error {
	if c.Previous == nil {
		return fmt.Errorf("you're on the first page")
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
