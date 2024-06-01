package updater

import (
	"fmt"
	"os"
	"strings"

	"github.com/gitarchived/service/internal/db"
	"github.com/gitarchived/service/internal/git"
	"github.com/gitarchived/service/internal/rabbit"
	"github.com/gitarchived/service/internal/s3"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Update(d *db.DB, s *s3.S3, a *amqp.Delivery) error {
	msg, err := rabbit.MessageRepositoryToJson(a.Body)

	if err != nil {
		return err
	}

	data := msg.Repository

	// Bundling the repo
	host, err := d.GetHostByName(data.Host)

	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://%s/%s/%s", host.Prefix, data.Owner, data.Name)

	err = git.Clone(url, data.Name)

	if err != nil {
		return err
	}

	err = git.Bundle(data.Name, data.ID)

	if err != nil {
		return err
	}

	path := splitPath(data.Name, data.ID)
	localPath := fmt.Sprintf("./%s", strings.Join(path, "/"))
	dir := strings.Join(path[:len(path)-1], "/")

	err = os.MkdirAll(dir, os.ModePerm)

	if err != nil {
		return err
	}

	err = os.Rename(fmt.Sprintf("./%s/%d.bundle", data.Name, data.ID), localPath)

	if err != nil {
		return err
	}

	// Update the last commit
	if err = d.UpdateRepositoryCommit(data.ID, msg.LastCommitKnown); err != nil {
		return err
	}

	// Uploading the bundle to S3
	// Why not return localPath? It's because S3 dosn't support ./ or ../ similiar symbols in front of the path
	// https://stackoverflow.com/questions/30518899/amazon-s3-how-to-fix-the-request-signature-we-calculated-does-not-match-the-s
	err = s.Upload(strings.Join(path, "/"))

	if err != nil {
		return err
	}

	// Remove the bundle from the local storage
	if err = Clear(data.Name); err != nil {
		return err
	}

	return nil
}

func splitPath(name string, id uint) []string {
	path := strings.Split(name, "")
	path = append(path, fmt.Sprintf("%d.bundle", id))

	for i, letter := range path {
		if letter == "." {
			path[i] = "-"
		}
	}

	return path
}
