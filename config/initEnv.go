package config

import (
	"log"
	"net"
	"net/url"
	"os"
)

var DB_URL string
var DB_HOST string
var DB_PORT string
var DB_USER string
var DB_PASSWORD string
var DB_NAME string

func InitEnv() {
	DB_URL = os.Getenv("DATABASE_URL")

	if DB_URL == "" {
		log.Print("There is no DATABASE_URL in env variables")
	}

	u, err := url.Parse(DB_URL)
	if err != nil {
		panic(err)
	}

	DB_HOST, DB_PORT, _ = net.SplitHostPort(u.Host)
	DB_USER = u.User.Username()
	DB_PASSWORD, _ = u.User.Password()
	DB_NAME = u.Path
}
