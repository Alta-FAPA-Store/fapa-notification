package main

import (
	"fmt"
	"go-hexagonal/email"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func newQueueConnection() *amqp.Connection {
	godotenv.Load(".env")
	// FROM := os.Getenv("FROM_EMAIL")
	// PASS_EMAIL := os.Getenv("EMAIL_PASS")
	RABBIT_USER := os.Getenv("RABBIT_USER")
	RABBIT_PASS := os.Getenv("RABBIT_PASS")
	ADDRESS_RABBIT := os.Getenv("ADDRESS_RABBIT")
	PORT_RABBIT := os.Getenv("PORT_RABBIT")

	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%s", RABBIT_USER, RABBIT_PASS, ADDRESS_RABBIT, PORT_RABBIT)

	conn, err := amqp.Dial(connectionString)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return conn
}

func main() {

	//NewConnection
	conn := newQueueConnection()
	email.ConsumerEmail(conn)

}
