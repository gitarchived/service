package main

import (
	"log"

	"github.com/gitarchived/service/internal/db"
	"github.com/gitarchived/service/internal/rabbit"
	"github.com/gitarchived/service/internal/s3"
	"github.com/gitarchived/service/internal/updater"
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

	s3, err := s3.Connect()

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

	qUpdate, err := ch.QueueDeclare(
		"update",
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
		qUpdate.Name, // queue
		"",           // consumer
		false,        // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)

	var forever chan struct{}

	go func() {
		for dl := range msgs {
			log.Printf("[+] Received message: %s\n", dl.Body)
			err := updater.Update(d, s3, &dl)

			if err != nil {
				log.Printf("[-] Error: %s\n", err)
			}

			dl.Ack(false)
		}
	}()

	log.Printf("[+] Waiting for messages")
	<-forever
}
