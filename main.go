package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-twitter-bot/config"
	"github.com/go-twitter-bot/modules/twitter"
)

var creds config.Config

func init() {
	filePath := config.GetPath("config.json")

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		config.CheckError("Error while reading file\n", err)
		os.Exit(1)
	}

	err = json.Unmarshal(file, &creds)
	if err != nil {
		log.Fatal("error:", err)
		os.Exit(1)
	}

}

func main() {
	fmt.Println("Go-Twitter Bot v0.01")
	client, err := twitter.CreateConnection(&creds.Credentials)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}

	messages := twitter.DMs(client)
	for _, event := range messages.Events {
		fmt.Printf("%+v\n", event)
		fmt.Printf("  %+v\n", event.Message)
		fmt.Printf("  %+v\n", event.Message.Data)
	}
}
