package main

import (
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

	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}
}
