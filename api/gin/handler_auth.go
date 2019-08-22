package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"

	myproto "github.com/gopperin/sme-ray/srv/auth/proto" // import proto生成的类
)

// Auth Auth
type Auth struct{}

var (
	clauth myproto.AuthService
)

func setupAuthClient() {
	clauth = myproto.NewAuthService("snc.gc.srv.auth", client.DefaultClient)
}

// Login Login
func (a *Auth) Login(c *gin.Context) {
	type Params struct {
		User string `json:"user"`
		Pwd  string `json:"pwd"`
	}

	var _params Params
	if err := c.ShouldBindJSON(&_params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_resp, err := clauth.Login(context.TODO(), &myproto.LoginRequest{
		User: _params.User, Pwd: _params.Pwd}, Filter("1.0.0"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, _resp)
}

// Verify Verify
func (a *Auth) Verify(c *gin.Context) {
	type Params struct {
		User string `json:"user"`
		Pwd  string `json:"pwd"`
	}

	var _params Params
	if err := c.ShouldBindJSON(&_params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_sub, _ := c.Get("sub")
	fmt.Println("sub:", _sub)

	// Set arbitrary headers in context
	_ctx := metadata.NewContext(context.Background(), map[string]string{"sub": _sub.(string)})

	_resp, err := clauth.VerifyToken(_ctx, &myproto.Request{}, Filter("1.0.0"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, _resp)
}
