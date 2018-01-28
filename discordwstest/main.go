package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/andersfylling/disgord/discordws"
	"github.com/sirupsen/logrus"
)

func main() {
	token := os.Getenv("DISGORD_TOKEN")
	if token == "" {
		panic("Missing disgord token in env var: DISGORD_TOKEN")
	}
	logrus.SetLevel(logrus.DebugLevel)
	termSignal := make(chan os.Signal, 1)
	signal.Notify(termSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	d := discordws.NewRequiredClient(&discordws.Config{
		Token:        token,
		DAPIVersion:  6,
		DAPIEncoding: discordws.EncodingJSON,
	})
	err := d.Connect()
	if err != nil {
		panic(err)
	}
	<-termSignal
	fmt.Println("Closing connection")
	err = d.Disconnect()
	if err != nil {
		logrus.Fatal(err)
	}
}