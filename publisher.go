package main

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func Publisher(ch *amqp.Channel) {

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
