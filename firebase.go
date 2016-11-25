package main

import (
	"github.com/treiff/ripta-transit-service/transit_realtime"
	"gopkg.in/zabawaba99/firego.v1"
	"log"
)

func updateDb(vehJson *transit_realtime.VehiclePosition, busId string) bool {
	url := "https://ri-realtime-transit.firebaseio.com/buses/" + busId
	f := firego.New(url, nil)
	f.Auth("AUTH_KEY")
	if err := f.Set(vehJson); err != nil {
		log.Fatal("set vehicle error: ", err)
		return false
	}
	return true
}
