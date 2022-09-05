package main

import (
	"fmt"
	"log"
	// "os"
	"encoding/json"
	// "net/http"
	// "time"
	// "io/ioutil"

	"github.com/bjp232004/go-freshdesk"
)

func main(){
	client, err := freshdesk.NewClient("zinccptest", "vTKUxPL7s25ou9qVOU3")
	if err != nil {
		log.Fatalf("Could not create client: %s", err)
	}


	// Create Ticket
	// ticket := &freshdesk.Ticket{
	// 	Priority : 1,
	// 	RequesterID : 88000825587,
	// 	ResponderID : 88000825587,
	// 	Source : 1,
	// 	Status : 2,
	// 	Subject : "Ticket Title3",
	// 	Name : "Ticket Title1",
	// 	Type : "Incident",
	// 	Email: "bjp232004@gmail.com",
	// 	Description : "this is a sample ticket",
	// }

	
	// if _, err := client.Tickets().Create(ticket); err != nil {
	// 	log.Fatalf("failed to create ticket: %s", err)
	// }
	fmt.Print("\n")
	
	
	// // List tickets
	resp, err := client.Tickets().ListAll();
	if err != nil {
		log.Fatalf("failed to create ticket: %s", err)
	}

	res2B, _ := json.Marshal(resp)

	var objmap []map[string]interface{}
	if err = json.Unmarshal(res2B, &objmap); err != nil {
		log.Fatal(err)
	}
	
	for _, value := range objmap {
		fmt.Print(value["id"],value["subject"])
		fmt.Print("\n")
	}
}