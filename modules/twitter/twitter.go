package twitter

import (
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/go-twitter-bot/config"
)

func CreateConnection(creds *config.Credentials) (*twitter.Client, error) {
	// Pass in your consumer key (API Key) and your Consumer Secret (API Secret)
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	// Pass in your Access Token and your Access Token Secret
	log.Println(creds.ConsumerKey)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	// we can retrieve the user and verify if the credentials
	// we have used successfully allow us to log in!
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	log.Printf("User's ACCOUNT:\n%+v\n", user)

	return client, nil

}

func Follow(user string, client twitter.Client) {
	_, _, err := client.Friendships.Create(&twitter.FriendshipCreateParams{
		ScreenName: user,
		Follow:     twitter.Bool(true),
	})
	config.LogError(err)
}

// List most recent 10 Direct Messages
func DMs(client *twitter.Client) *twitter.DirectMessageEvents {
	messages, _, err := client.DirectMessages.EventsList(
		&twitter.DirectMessageEventsListParams{Count: 10},
	)
	fmt.Println("User's DIRECT MESSAGES:")
	if err != nil {
		log.Fatal(err)
	}

	return messages
}
