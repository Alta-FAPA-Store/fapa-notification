package email

import (
	"fmt"

	"github.com/streadway/amqp"
)

func ConsumerEmail(conn *amqp.Connection) {
	ch, err := conn.Channel()
	defer conn.Close()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()

	msg, err := ch.Consume(
		"EmailQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range msg {
			fmt.Printf("Recieved Message: %s\n", d.Body)
			SendEmail(d.Body)
		}
	}()

	fmt.Println("Succesfully connected to our RabbitMQ instance")
	fmt.Println(" [*] - waiting for message")

	<-forever
}
