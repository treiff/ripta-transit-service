package main

import (
	"encoding/json"
	"fmt"
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
		//vehId := fe.GetVehicle().Vehicle.GetId()
		//lat := fe.GetVehicle().Position.GetLatitude()
		//long := fe.GetVehicle().Position.GetLongitude()
		//bearing := fe.GetVehicle().Position.GetBearing()
		//speed := fe.GetVehicle().Position.GetSpeed()

		//vehicle := Bus{vehId: Position{
		//	PositionFields{
		//		Latitude:  lat,
		//		Longitude: long,
		//		Speed:     speed,
		//		Bearing:   bearing,
		//	},
		//},
		//}

		//js, _ := json.MarshalIndent(vehicle, "", "  ")
		js, _ := json.MarshalIndent(fe.GetVehicle(), "", "  ")
		fmt.Printf("%s", js)
	}
}
