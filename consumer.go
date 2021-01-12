package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func Consumer(msgs <-chan amqp.Delivery) {

	for d := range msgs {
		fmt.Printf("Received: %s\n", string(d.Body))
	}

}
