package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"

	myproto "github.com/gopperin/sme-ray/srv/event/proto" // import proto生成的类
)

// Event Event
type Event struct{}

var (
	clEvent myproto.EventService
)

func setupEventClient() {
	clEvent = myproto.NewEventService("gc.micro.srv.pubsub.event", client.DefaultClient)
}

// SendEmail SendEmail
func (e *Event) SendEmail(c *gin.Context) {

	_resp, err := clEvent.SendEvent(context.TODO(),
		&myproto.Request{Id: "1", Message: "ministor@126.com"},
		Filter("1.0.0"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, _resp.Msg)
}
