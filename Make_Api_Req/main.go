package main

import (
	"Make_Api_Req/get_data"
	"Make_Api_Req/load_data"
	"Make_Api_Req/transform"
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	var input = &get_data.Input{NumberOfRecords: 50, Link: "https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=raw&sole"}
	users, err := get_data.AppendUsers(input, input.NumberOfRecords, ctx)
	if err != nil {
		log.Print("Error trying to get user data, ", err)
		return
	}
	users = transform.RemoveDuplicates(users)
	grouped_users := transform.Group_DataByIndex(users)
	var w = &load_data.WriteFile{}
	err = load_data.AppendRecords(grouped_users, w)
	if err != nil {
		log.Print("Error trying to load the data into json files")
		return
	}
	//log.Print(users)
}
