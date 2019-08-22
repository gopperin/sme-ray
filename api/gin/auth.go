package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

// APIJSONAuth APIJSONAuth
func APIJSONAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 把request的内容读取出来
		var _bodyBytes []byte
		if c.Request.Body == nil {
			APIAbortWithError(c, http.StatusOK, "Request.Body is Nil")
			return
		}

		_bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		_reader := bytes.NewReader(_bodyBytes)
		var _props map[string]interface{}
		err := BindJSON(_reader, &_props)
		if err != nil {
			APIAbortWithError(c, http.StatusOK, "BindJSON Error")
			return
		}

		_sign := strings.ToLower(APICalcSign(_props, GConfig.APIAppendKey, GConfig.APIMd5Key))
		fmt.Println("====== api signed : ", _sign)
		if _props[GConfig.APIMd5Key] != _sign {
			APIAbortWithError(c, http.StatusOK, "Invaild Signed")
			return
		}
		// 把刚刚读出来的再写进去
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(_bodyBytes))
		c.Next()
	}
}

// APIAbortWithError APIAbortWithError
func APIAbortWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"flag": 2,
		"msg":  message,
		"data": "",
	})
	c.Abort()
}

// JWTAuth JWTAuth
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		_token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(GConfig.JWTSign))
			return b, nil
		})

		if err != nil {
			JWTAbortWithError(c, http.StatusUnauthorized, "Invaild User Token", GConfig.JWTRealm)
			return
		}

		claims := _token.Claims.(jwt.MapClaims)

		c.Set("sub", claims["sub"])

		c.Next()
	}
}

// JWTAbortWithError JWTAbortWithError
func JWTAbortWithError(c *gin.Context, code int, message, realm string) {
	c.Header("WWW-Authenticate", "JWT realm="+realm)
	c.JSON(code, gin.H{
		"flag": code,
		"msg":  message,
	})
	c.Abort()
}
