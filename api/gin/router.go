package main

import (
	"github.com/gin-gonic/gin"
)

// SetupVersionRouter SetupVersionRouter
func SetupVersionRouter(g *gin.Engine) {
	_version := new(Version)
	_rversion := g.Group("/api/version")
	_rversion.Use(APIJSONAuth())
	{
		_rversion.POST("/info", _version.GetVersion)
		_rversion.POST("/check", _version.CheckVersion)
	}
}

// SetupAuthRouter SetupAuthRouter
func SetupAuthRouter(g *gin.Engine) {
	_auth := new(Auth)
	_r0 := g.Group("/api/auth")
	_r0.Use(APIJSONAuth())
	{
		_r0.POST("/login", _auth.Login)
	}

	_r1 := g.Group("/api/auth")
	_r1.Use(JWTAuth())
	_r1.Use(APIJSONAuth())
	{
		_r1.POST("/verify", _auth.Verify)
	}
}

// SetupEventRouter SetupEventRouter
func SetupEventRouter(g *gin.Engine) {
	_event := new(Event)
	_revent := g.Group("/api/event")
	_revent.Use(APIJSONAuth())
	{
		_revent.POST("/email", _event.SendEmail)
	}
}
