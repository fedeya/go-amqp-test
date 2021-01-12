package main

import (
	"github.com/streadway/amqp"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")

	handleError(err)

	defer conn.Close()

	ch, err := conn.Channel()

	handleError(err)

	defer ch.Close()

	q, err := ch.QueueDeclare("test", false, false, false, false, nil)

	handleError(err)

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)

	handleError(err)

	forever := make(chan bool)

	go Publisher(ch)
	go Consumer(msgs)

	<-forever

}
