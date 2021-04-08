package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/nstoker/go_tdd_tutorial/app"
	"github.com/sirupsen/logrus"
)

func main() {
	godotenv.Load(".env")
	a := app.App{}
	a.Initialize(os.Getenv("DATABASE_URL"))
	port := os.Getenv("PORT")
	if port == "" {
		logrus.Fatal("port environment variable missing")
	}
	a.Run(port)
}
