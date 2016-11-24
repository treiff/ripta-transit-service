package main

import (
	"io/ioutil"
	"net/http"
)

type Bus map[string]Position

type Position struct {
	PositionFields PositionFields `json:"position"`
}

type PositionFields struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Speed     float32 `json:"speed"`
	Bearing   float32 `json:"bearing"`
}

func fetchVehicleFeed() ([]byte, error) {
	url := "http://realtime.ripta.com:81/api/vehiclepositions?format=gtfs.proto"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
