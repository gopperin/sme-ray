package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	_ "github.com/micro/go-plugins/registry/etcdv3"

	"github.com/gopperin/sme-ray/srv/auth/proto"
)

// Auth Auth
type Auth struct {
}

// Login Login
func (s *Auth) Login(ctx context.Context, req *auth.LoginRequest, rsp *auth.LoginResult) error {

	fmt.Println(req.User, req.Pwd)

	_expire := time.Now().Add(time.Minute * time.Duration(GConfig.Expire))

	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["sub"] = req.User
	claims["exp"] = _expire.Unix()
	claims["iat"] = time.Now().Unix()
	claims["phone"] = "13810101010"
	claims["email"] = "ministor@126.com"
	token.Claims = claims
	// Sign and get the complete encoded token as a string
	_token, err := token.SignedString([]byte(GConfig.Sign))
	if err != nil {
		rsp.Token = "error"
		return nil
	}

	rsp.Token = _token

	return nil
}

// VerifyToken returns a customer from authentication token.
func (s *Auth) VerifyToken(ctx context.Context, req *auth.Request, rsp *auth.Result) error {
	md, _ := metadata.FromContext(ctx)

	fmt.Println(md)

	_sub := md["Sub"]

	rsp.State = 1
	rsp.Msg = _sub
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("snc.gc.srv.auth"),
		micro.Version("1.0.0"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)

	service.Init()

	auth.RegisterAuthHandler(service.Server(), &Auth{})

	service.Run()
}
