package get_data

import (
	"context"
	"encoding/json"
)

type User struct {
	First   string `json:"first"`
	Last    string `json:"last"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Created string `json:"created"`
	Balance string `json:"balance"`
}

func AppendUsers(read Extract, numberExtracted int, ctx context.Context) ([]User, error) {
	var users []User
	client := sendRequest{}
	for {
		var tempData []User
		body, err := read.extractData(client, ctx)
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(body, &tempData); err != nil {
			return nil, err
		}
		users = append(users, tempData...)
		if len(users) >= numberExtracted {
			//discard users after the element var is passed
			users = users[:numberExtracted]
			break
		}
	}
	return users, nil
}
