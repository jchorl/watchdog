package client

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"

	pb "github.com/jchorl/watchdog/proto"
)

// Client is a client to talk to watchdog
type Client struct {
	Domain string
}

// Ping pings the server
func (c Client) Ping(name string, frequency pb.Watch_Frequency) error {
	watch := &pb.Watch{
		Name:      name,
		Frequency: frequency,
	}

	data, err := proto.Marshal(watch)
	if err != nil {
		log.Fatal("Error marshaling the watch: ", err)
		return err
	}

	resp, err := http.Post(fmt.Sprintf("%s/ping", c.Domain), "application/x-protobuf", bytes.NewReader(data))
	if err != nil {
		log.Fatal("Error posting: ", err)
		return err
	}
	if resp.StatusCode != 200 {
		log.Fatal("Ping did not return 200")
		return errors.New("Ping did not return 200")
	}

	return nil
}
