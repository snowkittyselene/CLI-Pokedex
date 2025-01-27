package pokeapi

type LocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

const locationAreaUrl = baseApiUrl + "location-area/"

func (client *Client) CallLocationArea(url *string) (LocationArea, error) {
	pageUrl := locationAreaUrl
	if url != nil {
		pageUrl = *url
	}

	data, err := CallAPI(client, pageUrl, LocationArea{})
	if err != nil {
		return LocationArea{}, err
	}
	return data, nil
}
