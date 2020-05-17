package main

import (
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/FusionAuth/go-client/pkg/fusionauth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const (
	clientID string = "8cc4c901-1852-4e62-a655-699f2c94ffdc"
	clientSecret string = "OshhDO0d1Y6Hrejsd697Yo4unTar_gI00fHiIASMZGc"
	host string = "http://localhost:9011"
	apiKey = "vB_ap-t-zsCbOYv9HhETbbRm1Ue8C3FFP28qJQWfNTo"
)

var (
	httpClient = &http.Client{
		Timeout: time.Second * 10,
	}
	baseURL, _ = url.Parse(host)
	client = fusionauth.NewClient(httpClient, baseURL, apiKey)
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("user_session", store))

	r.LoadHTMLGlob("templates/*")

	r.GET("/", indexRoute)
	r.GET("/oauth/redirect", oauthRedirectRoute)
	r.GET("/authenticated", authdRoute)

	return r
}

func indexRoute(c *gin.Context) {
	session := sessions.Default(c)
	userFirstName, ok := session.Get("userFirstName").(string)

	name := ""
	if ok && userFirstName != "" {
		name = userFirstName
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{ "Name": name, "ClientID": clientID })
}

func oauthRedirectRoute(c *gin.Context) {
	token, oauthErr, err := client.ExchangeOAuthCodeForAccessToken(
		c.Query("code"),
		clientID,
		clientSecret,
		"http://localhost:8080/oauth/redirect",
	)
	if err != nil {
		log.Fatal(err)
		c.String(http.StatusInternalServerError, "There was an issue.")
	}

	log.Printf("oauthError: %+v", oauthErr)
	log.Printf("token: %+v", token.AccessToken)

	userResp, faErrors, err := client.RetrieveUserUsingJWT(token.AccessToken)
	if err != nil {
		log.Fatal(err)
		c.String(http.StatusInternalServerError, "There was an issue.")
	}

	log.Printf("faError: %+v", faErrors)
	log.Printf("user: %+v", userResp.User)

	session := sessions.Default(c)
	session.Set("userFirstName", userResp.User.FirstName)
	session.Save()

	c.Redirect(http.StatusMovedPermanently, "/")
}

func authdRoute(c *gin.Context) {
	c.String(http.StatusOK, "You're Auth'd!")
}