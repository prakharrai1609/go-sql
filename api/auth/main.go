package auth

import (
	"prakharrai1609/config"
	"prakharrai1609/db/queries"
	"prakharrai1609/models"
	"prakharrai1609/response"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	db := config.Get().Db

	name, mail := c.Query("name"), c.Query("mail")
	queryString := queries.FetchOneUser(name, mail)

	res := db.QueryRow(queryString)

	var user models.Users

	err := res.Scan(&user.Id, &user.Name, &user.Mail)

	if err != nil {
		c.JSON(500, response.Error("Invalid user, first sign in."))
		return
	}

	c.JSON(200, response.Response("Valid user."))
}
