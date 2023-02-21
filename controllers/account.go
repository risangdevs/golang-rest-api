package controllers

import (
	"golang-rest-api/db"
	"golang-rest-api/repositories"
	"golang-rest-api/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllAccount(c *gin.Context) {
	var (
		result gin.H
	)

	accounts, err := repositories.GetAllAccount(db.DbConnection)
	if err != nil {
		result = gin.H{"result": err}
	} else {
		result = gin.H{"result": accounts}
	}
	c.JSON(http.StatusOK, result)

}

func InsertAccount(c *gin.Context) {
	var account structs.Account

	err := c.ShouldBindJSON(&account)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"result": "Success Insert Account"})
}

func UpdateAccount(c *gin.Context) {
	var account structs.Account
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&account)
	if err != nil {
		panic(err)
	}
	account.ID = int64(id)
	err = repositories.UpdateAccount(db.DbConnection, account)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"result": "Success Update Account"})

}

func DeleteAccount(c *gin.Context) {
	var account structs.Account
	id, err := strconv.Atoi(c.Param("id"))

	account.ID = int64(id)

	err = repositories.DeleteAccount(db.DbConnection, account)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"result": "Success Delete Account"})
}
