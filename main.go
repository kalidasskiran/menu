package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Account struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Username     string    `json:"Username"`
	Password     string    `json:"password"`
	Timeofsignup time.Time `json:"timeofsignup"`
}

var accounts []Account

func init() {
	accounts = make([]Account, 0)
}

func NewAccountHandler(c *gin.Context) {
	var account Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	account.ID = xid.New().String()
	account.Timeofsignup = time.Now()
	accounts = append(accounts, account)
	c.JSON(http.StatusOK, account)
}

func main() {
	router := gin.Default()
	router.POST("/signup", NewAccountHandler)
	router.Run()
}
