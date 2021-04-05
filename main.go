package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/constantgillet/resthelpdesk/config"
	"github.com/constantgillet/resthelpdesk/routes"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	//Init env variables
	config.InitEnv()

	PORT, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		PORT = 8000
	}

	http.Handle("/", routes.Handlers())
	fmt.Println("App started on ", PORT)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), nil))
}
