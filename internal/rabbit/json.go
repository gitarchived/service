package rabbit

import "encoding/json"

func MessageRepositoryToJson(body []byte) (Repository, error) {
	var msg Repository

	err := json.Unmarshal(body, &msg)

	return msg, err
}
