package main

import (
	"fmt"
	"net/http"
	"prakharrai1609/config"

	"prakharrai1609/api/auth"
	"prakharrai1609/api/unauth"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	unauthenticated := router.Group("")
	{
		unauthenticated.GET("healthcheck", unauth.Healthcheck)
		unauthenticated.POST("login", unauth.Login)
	}

	authenticated := router.Group("auth")
	{
		authenticated.POST("signin", auth.SignIn)
	}

	PORT := fmt.Sprintf("localhost:%s", config.Get().Port)

	fmt.Printf("Server is up and running on port %s ✨\n\n", PORT)

	err := http.ListenAndServe(PORT, router)

	defer config.Get().Db.Close()

	if err != nil {
		panic(err)
	}
}
