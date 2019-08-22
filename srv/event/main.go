package main

import (
	"context"
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
	proto "github.com/gopperin/sme-ray/srv/event/proto"
)

const topic = "gc.topic.pubsub.email"

var mybroker broker.Broker

// Event Event
type Event struct {
}

// SendEvent SendEvent
func (s *Event) SendEvent(ctx context.Context, req *proto.Request, rsp *proto.Result) error {
	var _email email.Email
	_email.To = req.Message

	_body, _ := json.Marshal(&_email)

	msg := broker.Message{
		Header: map[string]string{
			"id": req.Id,
		},
		Body: _body,
	}
	err := mybroker.Publish(topic, &msg)
	if err != nil {
		fmt.Println(err.Error())
	}

	rsp.Msg = "Send email ok"
	return nil
}

func main() {
	// create a service
	service := micro.NewService(
		micro.Name("gc.micro.srv.pubsub.event"),
		micro.Version("1.0.0"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)
	// parse command line
	service.Init()

	// Get the broker instance using our environment variables
	mybroker = service.Server().Options().Broker
	if err := mybroker.Connect(); err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}

	// 注册服务
	proto.RegisterEventHandler(service.Server(), new(Event))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
