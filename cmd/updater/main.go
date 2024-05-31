package main

import (
	"encoding/json"
	"log"

	"github.com/gitarchived/service/internal/db"
	"github.com/gitarchived/service/internal/rabbit"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	_, err = db.Connect()

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
			var data db.Repository
			err := json.Unmarshal(dl.Body, &data)

			if err != nil {
				log.Printf("[-] Error: %v", err)
				continue
			}

			log.Printf("[+] Processing: %v", data.ID)
			dl.Ack(false)
		}
	}()

	log.Printf("[+] Waiting for messages")
	<-forever
}
