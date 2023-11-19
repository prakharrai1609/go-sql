package unauth

import (
	"fmt"
	"log"
	"prakharrai1609/config"
	"prakharrai1609/db/queries"
	"prakharrai1609/response"

	"github.com/gin-gonic/gin"
)

func Healthcheck(c *gin.Context) {
	c.JSON(200, response.Response("Auth service is healthy"))
}

func Login(c *gin.Context) {
	name, mail := c.Query("name"), c.Query("mail")

	queryString := queries.CreateUser(name, mail)

	fmt.Println(name, mail, queryString)

	_, err := config.Get().Db.Exec(queryString)

	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(200, response.Response("User created."))
}
