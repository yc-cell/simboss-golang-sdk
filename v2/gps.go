package simboss

import (
	"encoding/json"
	"net/url"
)

type LocationService struct {
	client *Client
}

type Location struct {
	Iccid      string `json:"iccid"`
	Longitude  string `json:"longitude"`
	Latitude   string `json:"latitude"`
	Coordinate string `json:"coordinate"`
	Location   string `json:"location"`
}

//定位查询

func (l *LocationService) LocationSearch(params url.Values) (*Location, error) {
	if err := RequiredCardId(params); err != nil {
		return nil, err
	}
	location := &Location{}
	body, err := l.client.Post("/lbs/location/search", params)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, location); err != nil {
		return nil, err
	}
	return location, nil
}
