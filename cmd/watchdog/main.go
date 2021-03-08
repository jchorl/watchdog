package main

import (
	"flag"

	"github.com/golang/glog"

	"github.com/jchorl/watchdog/pkg/client"
	pb "github.com/jchorl/watchdog/proto"
)

var (
	name         = flag.String("name", "", "name of the service")
	frequencyStr = flag.String("frequency", "", "frequency to update, either daily or weekly")
	domain       = flag.String("domain", "", "watchdog domain")
)

func main() {
	flag.Set("logtostderr", "true") // for glog
	flag.Parse()

	if *name == "" {
		glog.Fatalf("name cannot be empty")
	}

	if *domain == "" {
		glog.Fatalf("domain cannot be empty")
	}

	var frequency pb.Watch_Frequency
	switch *frequencyStr {
	case "daily":
		frequency = pb.Watch_DAILY
	case "weekly":
		frequency = pb.Watch_WEEKLY
	default:
		glog.Fatalf("frequency must be daily or weekly, got: %s", *frequencyStr)
	}

	c := &client.Client{
		Domain: *domain,
	}

	err := c.Ping(*name, frequency)
	if err != nil {
		glog.Fatalf("ping failed: %+v", err)
	}

	glog.Infof("ping succeeded for %s", *name)
}
