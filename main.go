package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/treiff/ripta-transit-service/transit_realtime"
	"log"
)

func main() {

	feed, err := fetchVehicleFeed()
	if err != nil {
		log.Fatal("fetch vehicle feed error: ", err)
	}

	fm := new(transit_realtime.FeedMessage)
	err = proto.Unmarshal(feed, fm)

	for _, fe := range fm.Entity {
		vehId := fe.GetVehicle().Vehicle.GetId()
		position := fe.GetVehicle().Position
		updatePosition(position, vehId)
	}
}
