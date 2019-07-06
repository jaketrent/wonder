package users

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ok struct {
	Data []*User `json:"data"`
}

type clienterr struct {
	Title  string `json:"title"`
	Status int    `json:"status"`
}

type bad struct {
	Errors []clienterr `json:"errors"`
}

func list(c *gin.Context) {
	db, _ := c.MustGet("db").(*sql.DB)
	users, err := findAll(db)

	if err == nil {
		c.JSON(http.StatusOK, ok{
			Data: users,
		})
	} else {
		c.JSON(http.StatusInternalServerError, bad{Errors: []clienterr{{Title: "Server error", Status: http.StatusInternalServerError}}})
		fmt.Println("users list error", err)
	}
}

func Mount(router *gin.Engine) {
	router.GET("/api/v1/users", list)
}
