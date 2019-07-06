package users

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		c.JSON(http.StatusOK, struct {
			Data []*User `json:"data"`
		}{
			Data: users,
		})

	} else {
		c.JSON(http.StatusInternalServerError, bad{Errors: []clienterr{{Title: "Server error", Status: http.StatusInternalServerError}}})
		fmt.Println("users list error", err)
	}
}

func createWonder(c *gin.Context) {
	db, _ := c.MustGet("db").(*sql.DB)

	var err error
	var req struct {
		Data []*UserWonder `json:"data"`
	}
	err = c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, bad{
			Errors: []clienterr{{Title: "Cannot create wonder for user", Status: http.StatusBadRequest}},
		})
		fmt.Println("user addWonder req error", err)
		return
	}
	userWonders, err := insertWonder(db, req.Data)

	if err == nil {
		c.JSON(http.StatusOK, struct {
			Data []*UserWonder `json:"data"`
		}{
			Data: userWonders,
		})
	} else {
		c.JSON(http.StatusInternalServerError, err)
		fmt.Println("user addWonder db error", err)
	}
}

// Mount connects routes to router
func Mount(router *gin.Engine) {
	router.GET("/api/v1/users", list)
	router.POST("/api/v1/users/wonders", createWonder)
}
