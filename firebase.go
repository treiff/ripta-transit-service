package main

import (
	"github.com/treiff/ripta-transit-service/transit_realtime"
	"gopkg.in/zabawaba99/firego.v1"
	"log"
)

func updatePosition(posJson *transit_realtime.Position, busId string) bool {
	url := "https://ri-realtime-transit.firebaseio.com/busposition/" + busId
	f := firego.New(url, nil)
	//f.Auth("AUTH_KEY")
	if err := f.Set(posJson); err != nil {
		log.Fatal("set vehicle error: ", err)
		return false
	}
	return true
}
