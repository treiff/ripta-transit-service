package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/treiff/transit_service/transit_realtime"
	"log"
)

func main() {
	test := &transit_realtime.VehicleDescriptor{
		Id:           proto.String("1"),
		Label:        proto.String("test vehicle"),
		LicensePlate: proto.String("12345"),
	}

	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	newTest := &transit_realtime.VehicleDescriptor{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}

	feed, err := fetchVehicleFeed()
	if err != nil {
		log.Fatal("fetch vehicle feed error: ", err)
	}

	fm := new(transit_realtime.FeedMessage)
	err = proto.Unmarshal(feed, fm)
	//fmt.Printf("%v", fm)
	fmt.Printf("%v", fm.GetEntity())

}
