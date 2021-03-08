package server

import (
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/golang/protobuf/proto"

	pb "github.com/jchorl/watchdog/proto"
)

// datastore struggles with protobufs.
// so we can just wrap Watch and implement PropertyLoadSaver
type serializedWatch struct {
	Serialized []byte
	Proto      *pb.Watch
}

func (x *serializedWatch) Load(ps []datastore.Property) error {
	if err := datastore.LoadStruct(x, ps); err != nil {
		return err
	}

	if err := proto.Unmarshal(x.Serialized, x.Proto); err != nil {
		return fmt.Errorf("unmarshaling proto: %w", err)
	}

	return nil
}

func (x *serializedWatch) Save() ([]datastore.Property, error) {
	var err error
	x.Serialized, err = proto.Marshal(x.Proto)
	if err != nil {
		return nil, fmt.Errorf("marshaling proto: %w", err)
	}

	savedProto := x.Proto
	x.Proto = nil
	props, err := datastore.SaveStruct(x)
	x.Proto = savedProto
	if err != nil {
		return props, err
	}

	return props, nil
}
