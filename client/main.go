package main

import (
	"bytes"
	"errors"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/jchorl/watchdog/types"
)

// Ping pings the server
func Ping(name string, frequency types.Watch_Frequency) error {
	watch := &types.Watch{
		Name:      name,
		Frequency: frequency,
	}

	data, err := proto.Marshal(watch)
	if err != nil {
		log.Fatal("Error marshaling the watch: ", err)
		return err
	}

	resp, err := http.Post("http://localhost:8080/ping", "application/x-protobuf", bytes.NewReader(data))
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

func main() {
	Ping("myservice", types.Watch_WEEKLY)
}
