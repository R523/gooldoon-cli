package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/plgd-dev/go-coap/v2/udp"
)

// ArgsNum is the number of required arguments for application.
const ArgsNum = 3

func main() {
	if len(os.Args) != ArgsNum {
		log.Printf("usage: coap-demo [address e.g. 192.168.73.192:1378] [path e.g. elahe]\n")

		return
	}

	co, err := udp.Dial(os.Args[1])
	if err != nil {
		log.Fatalf("error dialing: %v", err)
	}

	path := os.Args[2]

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	resp, err := co.Get(ctx, path)
	if err != nil {
		log.Printf("error sending request: %v", err)

		return
	}

	log.Printf("response code: %s", resp.Code())

	body, err := resp.ReadBody()
	if err != nil {
		log.Printf("cannot read the payload %v", err)

		return
	}

	var r map[string]interface{}
	if err := json.Unmarshal(body, &r); err != nil {
		log.Printf("cannot parse the payload as json %v", err)

		return
	}

	log.Printf("response payload: %+v", r)
}
