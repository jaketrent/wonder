package users

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Mount connects routes to router
func Mount(router *gin.Engine) {
	router.GET("/api/v1/users", list)
	router.GET("/api/v1/users/:userId/wonders", listWonders)
	router.POST("/api/v1/users/:userId/wonders", createWonder)
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
		c.JSON(http.StatusOK, struct {
			Data []*User `json:"data"`
		}{
			Data: users,
		})

	} else {
		c.JSON(http.StatusInternalServerError, bad{Errors: []clienterr{{Title: "Cannot retrieve users", Status: http.StatusInternalServerError}}})
		fmt.Println("users list error", err)
	}
}

func listWonders(c *gin.Context) {
	db, _ := c.MustGet("db").(*sql.DB)

	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, bad{
			Errors: []clienterr{{Title: "Bad userId", Status: http.StatusBadRequest}},
		})
		return
	}

	userWonders, err := findWonders(db, userID)

	if err == nil {
		c.JSON(http.StatusOK, struct {
			Data []*UserWonder `json:"data"`
		}{
			Data: userWonders,
		})

	} else {
		c.JSON(http.StatusInternalServerError, bad{Errors: []clienterr{{Title: "Cannot retrieve user wonders", Status: http.StatusInternalServerError}}})
		fmt.Println("user wonders list error", err)
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
			Errors: []clienterr{{Title: "Malformed user wonder", Status: http.StatusBadRequest}},
		})
		fmt.Println("malformed user addWonder req error", err)
		return
	}

	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, bad{
			Errors: []clienterr{{Title: "Bad userId", Status: http.StatusBadRequest}},
		})
		return
	}

	userWonders, err := insertWonder(db, userID, req.Data)

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
