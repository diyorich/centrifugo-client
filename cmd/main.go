package main

import (
	"centrifugo-client-express24/internal"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Starting client application")
	if err := run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Application stopped")
}

func run() error {
	cfgPath, err := internal.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := internal.NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	closeClientConn, err := internal.Start(*cfg)
	defer closeClientConn()

	if err != nil {
		log.Fatal(err)
	}

	select {}

	return nil
}
