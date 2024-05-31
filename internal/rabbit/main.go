package rabbit

import (
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	*amqp.Connection
}

func Connect() (Rabbit, error) {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))

	if err != nil {
		return Rabbit{}, err
	}

	return Rabbit{conn}, nil
}
