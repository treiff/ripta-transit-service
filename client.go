package main

import (
	"io/ioutil"
	"net/http"
)

func fetchVehicleFeed() ([]byte, error) {
	url := "http://realtime.ripta.com:81/api/vehiclepositions?format=gtfs.proto"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
