package internal

import (
	"github.com/centrifugal/centrifuge-go"
	"log"
)

func Start(config AppConfig) (func(), error) {
	centrifugoConfig := centrifuge.Config{
		Token: config.Token,
	}

	client := centrifuge.NewProtobufClient(config.Endpoint, centrifugoConfig)

	client.OnConnecting(func(event centrifuge.ConnectingEvent) {
		log.Println("Connecting to centrifugo ...")
	})

	client.OnConnected(func(event centrifuge.ConnectedEvent) {
		log.Println("CONNECTED!")
		log.Printf("Connected to centrifugo with id %v", event.ClientID)
	})

	client.OnDisconnected(func(event centrifuge.DisconnectedEvent) {
		log.Println("Client disconnected")
	})

	client.OnError(func(e centrifuge.ErrorEvent) {
		log.Printf("Error: %s", e.Error.Error())
	})

	client.OnPublication(func(event centrifuge.ServerPublicationEvent) {
		log.Println("New Publication event")
		log.Printf("data %v", string(event.Data))
	})

	client.OnSubscribing(func(e centrifuge.ServerSubscribingEvent) {
		log.Printf("Subscribing on server side channel %s", e.Channel)
	})

	err := client.Connect()
	if err != nil {
		return nil, err
	}

	return client.Close, nil
}
