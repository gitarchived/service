package main

import (
	"log"

	"github.com/gitarchived/service/internal/db"
	"github.com/gitarchived/service/internal/deleter"
	"github.com/gitarchived/service/internal/rabbit"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	d, err := db.Connect()

	if err != nil {
		log.Fatal(err)
	}

	conn, err := rabbit.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Fatal(err)
	}

	qDelete, err := ch.QueueDeclare(
		"delete",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	// Consume
	msgs, err := ch.Consume(
		qDelete.Name, // queue
		"",           // consumer
		false,        // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)

	if err != nil {
		log.Fatal(err)
	}

	var forever chan struct{}

	go func() {
		for dl := range msgs {
			err := deleter.Delete(d, &dl)

			if err != nil {
				log.Printf("[-] Error delete: %s", err)
			}

			err = dl.Ack(false)

			if err != nil {
				log.Printf("[-] Error ack: %s", err)
			}
		}
	}()

	log.Printf("[+] Waiting for messages")
	<-forever
}
