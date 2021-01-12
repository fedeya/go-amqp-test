package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

func Publisher(ch *amqp.Channel, wg *sync.WaitGroup) {

	defer wg.Done()

	i := 0

	for {
		time.Sleep(time.Millisecond * 500)

		body := fmt.Sprintf("hello bro: %v", i)

		err := ch.Publish("", "test", false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

		fmt.Println("Sended: ", body)

		handleError(err)

		i++
	}

}
