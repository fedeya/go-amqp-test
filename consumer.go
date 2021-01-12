package main

import (
	"fmt"
	"sync"

	"github.com/streadway/amqp"
)

func Consumer(msgs <-chan amqp.Delivery, wg *sync.WaitGroup) {

	defer wg.Done()

	for d := range msgs {
		fmt.Printf("Received: %s\n", string(d.Body))
	}

}
