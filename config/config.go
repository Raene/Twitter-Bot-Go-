package config

import (
	"log" 
	"os"
)

// Config holds configuration from the user
type Config struct {
	TwitterName string
	Interests   []string
	Credentials Credentials
}

// Credentials stores all of our access/consumer tokens
// and secret keys needed for authentication against
// the twitter REST API.
type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func GetPath(filename string) string {
	pwd, _ := os.Getwd()
	return pwd + string(os.PathSeparator) + filename
	// _, filename, _, _ := runtime.Caller(0)
	// baseDir := filepath.Join(filename + "/../")
	// log.Println(baseDir + string(os.PathSeparator)+ "config.json")
}

func CheckError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func LogError(err error) {
	if err != nil {
		log.Println(err)
	}
}
