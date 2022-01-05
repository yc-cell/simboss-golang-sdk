package simboss

import (
	"encoding/json"
	"github.com/yc-cell/simboss-golang-sdk/utils"
	"github.com/yc-cell/simboss-golang-sdk/utils/time"
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

type Locations struct {
	TradeNo      string    `json:"tradeNo"`
	Iccid        string    `json:"iccid"`
	Longitude    string    `json:"longitude"`
	Latitude     string    `json:"latitude"`
	Coordinate   string    `json:"coordinate"`
	Location     string    `json:"location"`
	QueryType    string    `json:"queryType"`
	Fee          float64   `json:"fee"`
	QueryDate    time.Time `json:"queryDate"`
	LocationDate time.Time `json:"locationDate"`
	Memo         string    `json:"memo"`
}

type LocationList struct {
	Page Page        `json:"page"`
	List []Locations `json:"list"`
}

//定位查询记录

func (l *LocationService) List(params url.Values) (*LocationList, error) {
	if err := RequiredCardId(params); err != nil {
		return nil, err
	}
	if !utils.Required(params, "pageNo") {
		return nil, ErrRequired
	}
	locationList := &LocationList{
		Page: Page{},
		List: make([]Locations, 0),
	}
	body, err := l.client.Post("/lbs/record/list", params)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, locationList); err != nil {
		return nil, err
	}
	return locationList, nil
}
