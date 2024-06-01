package deleter

import (
	"fmt"

	"github.com/gitarchived/service/internal/db"
	"github.com/gitarchived/service/internal/git"
	"github.com/gitarchived/service/internal/rabbit"
	"github.com/gitarchived/service/internal/util"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Delete(d *db.DB, a *amqp.Delivery) error {
	msg, err := rabbit.MessageRepositoryToJson(a.Body)

	if err != nil {
		return err
	}

	data := msg.Repository

	// Host connettivity check (if host is reachable the repository will not be deleted)
	host, err := d.GetHostByName(data.Host)

	if err != nil {
		return err
	}

	if !util.IsOk(host.URL) {
		return fmt.Errorf("Host %s is unreachable", data.Host)
	}

	// Check if the repository exists
	if !util.IsOk(fmt.Sprintf("https://%s/%s/%s", host.Prefix, data.Owner, data.Name)) {
		if err := d.DeleteRepository(data.ID); err != nil {
			return err
		}
	}

	if _, err := git.RemoteLastCommit(fmt.Sprintf("https://%s/%s/%s.git", host.Prefix, data.Owner, data.Name)); err != nil {
		if err := d.DeleteRepository(data.ID); err != nil {
			return err
		}
	}

	return nil // Repository is valid and exists
}
