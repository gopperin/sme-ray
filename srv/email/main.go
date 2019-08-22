package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/broker/nsq"
	_ "github.com/micro/go-plugins/registry/etcdv3"

	email "github.com/gopperin/sme-ray/srv/email/proto"
)

const topic = "gc.topic.pubsub.email"

func main() {
	// create a service
	service := micro.NewService(
		micro.Name("gc.micro.srv.pubsub.email"),
		micro.Version("1.0.0"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)
	// parse command line
	service.Init()

	// Get the broker instance using our environment variables
	pubsub := service.Server().Options().Broker
	if err := pubsub.Connect(); err != nil {
		log.Fatal(err)
	}

	// Subscribe to messages on the broker
	_, err := pubsub.Subscribe(topic, func(p broker.Publication) error {
		var _email *email.Email
		if err := json.Unmarshal(p.Message().Body, &_email); err != nil {
			return err
		}

		go sendEmail(_email)
		return nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func sendEmail(email *email.Email) error {
	fmt.Println("success sending email to:", email.To)
	return nil
}
