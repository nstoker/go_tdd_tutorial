package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	godotenv.Load(".env")
	a := App{}
	a.Initialize(os.Getenv("DATABASE_URL"))
	port := os.Getenv("PORT")
	if port == "" {
		logrus.Fatal("port environment variable missing")
	}
	a.Run(port)
}
