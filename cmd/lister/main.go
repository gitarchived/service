package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gitarchived/service/internal/db"
	"github.com/gitarchived/service/internal/git"
	"github.com/gitarchived/service/internal/rabbit"
	"github.com/go-co-op/gocron/v2"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
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

	qRemove, err := ch.QueueDeclare(
		"remove",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	s, err := gocron.NewScheduler()

	if err != nil {
		log.Fatal(err)
	}

	_, err = s.NewJob(
		gocron.DurationJob(time.Hour),
		gocron.NewTask(
			func() {
				var repos []db.Repository

				if err := d.Find(&repos).Error; err != nil {
					log.Println(err)
				}

				for _, repo := range repos {
					host, err := d.GetHostByName(repo.Host)

					if err != nil {
						log.Println(err)
					}

					url := fmt.Sprintf("https://%s/%s/%s.git", host.Prefix, repo.Owner, repo.Name)
					commit, err := git.RemoteLastCommit(url)

					if err != nil {
						json, err := json.Marshal(repo)

						if err != nil {
							log.Println(err)
						}

						err = ch.PublishWithContext(
							ctx,
							"",
							qRemove.Name,
							false,
							false,
							amqp.Publishing{
								ContentType: "application/json",
								Body:        []byte(json),
							},
						)
					}

					if repo.LastCommit != commit {
						json, err := json.Marshal(rabbit.Repository{
							Repository:      repo,
							LastCommitKnown: commit,
						})

						if err != nil {
							log.Println(err)
						}

						err = ch.PublishWithContext(
							ctx,
							"",
							qUpdate.Name,
							false,
							false,
							amqp.Publishing{
								ContentType: "application/json",
								Body:        []byte(json),
							},
						)
					}
				}
			},
		),
	)

	if err != nil {
		fmt.Println(err)
	}

	s.Start()

	select {} // Block forever
}
