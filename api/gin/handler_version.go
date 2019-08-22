package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"

	myproto "github.com/gopperin/sme-ray/srv/version/proto" // import proto生成的类

	"context"
)

// Version Version
type Version struct{}

var (
	cl myproto.VersionService
)

func setupVersionClient() {
	cl = myproto.NewVersionService("snc.gc.srv.version", client.DefaultClient)
}

// GetVersion GetVersion
func (s *Version) GetVersion(c *gin.Context) {
	type Params struct {
		Version string `json:"version"`
	}

	var _params Params
	if err := c.ShouldBindJSON(&_params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := cl.GetVersion(
		context.TODO(),
		&myproto.VersionRequest{Version: _params.Version},
		Filter("1.0.0"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// CheckVersion CheckVersion
func (s *Version) CheckVersion(c *gin.Context) {
	type Params struct {
		Version string `json:"version"`
	}

	var _params Params
	if err := c.ShouldBindJSON(&_params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := cl.CheckVersion(context.TODO(), &myproto.CheckVersionRequest{
		Version: _params.Version,
	}, Filter("1.0.0"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
